// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

import "./Seerlink.sol";
import "./interfaces/ENSInterface.sol";
import "./interfaces/LinkTokenInterface.sol";
import "./interfaces/SeerlinkRequestInterface.sol";
import "./interfaces/PointerInterface.sol";
import { ENSResolver as ENSResolver_Seerlink } from "./vendor/ENSResolver.sol";

/**
 * @title The SeerlinkClient contract
 * @notice Contract writers can inherit this contract in order to create requests for the
 * Seerlink network
 */
contract SeerlinkClient {
  using Seerlink for Seerlink.Request;

  uint256 constant internal LINK_DIVISIBILITY = 10**18;
  uint256 constant private AMOUNT_OVERRIDE = 0;
  address constant private SENDER_OVERRIDE = address(0);
  uint256 constant private ARGS_VERSION = 2;
  bytes32 constant private ENS_TOKEN_SUBNAME = keccak256("link");
  bytes32 constant private ENS_ORACLE_SUBNAME = keccak256("oracle");
  address constant private LINK_TOKEN_POINTER = 0xA68ce5559cA8eACaA9f3CB308CC77a6E6188713b;

  ENSInterface private ens;
  bytes32 private ensNode;
  LinkTokenInterface private link;
  SeerlinkRequestInterface private oracle;
  uint256 private requestCount = 1;
  mapping(bytes32 => address) private pendingRequests;

  event SeerlinkRequested(bytes32 indexed id);
  event SeerlinkFulfilled(bytes32 indexed id);
  event SeerlinkCancelled(bytes32 indexed id);

  /**
   * @notice Creates a request that can hold additional parameters
   * @param specId The Job Specification ID that the request will be created for
   * @param callbackAddress The callback address that the response will be sent to
   * @param callbackFunctionSignature The callback function signature to use for the callback address
   * @return A Seerlink Request struct in memory
   */
  function buildSeerlinkRequest(
    bytes32 specId,
    address callbackAddress,
    bytes4 callbackFunctionSignature
  ) internal pure returns (Seerlink.Request memory) {
    Seerlink.Request memory req;
    return req.initialize(specId, callbackAddress, callbackFunctionSignature);
  }

  /**
   * @notice Creates a Seerlink request to the stored oracle address
   * @dev Calls `seerlinkRequestTo` with the stored oracle address
   * @param req The initialized Seerlink Request
   * @param payment The amount of LINK to send for the request
   * @return requestId The request ID
   */
  function sendSeerlinkRequest(Seerlink.Request memory req, uint256 payment)
    internal
    returns (bytes32)
  {
    return sendSeerlinkRequestTo(address(oracle), req, payment);
  }

  /**
   * @notice Creates a Seerlink request to the specified oracle address
   * @dev Generates and stores a request ID, increments the local nonce, and uses `transferAndCall` to
   * send LINK which creates a request on the target oracle contract.
   * Emits SeerlinkRequested event.
   * @param oracleAddress The address of the oracle for the request
   * @param req The initialized Seerlink Request
   * @param payment The amount of LINK to send for the request
   * @return requestId The request ID
   */
  function sendSeerlinkRequestTo(address oracleAddress, Seerlink.Request memory req, uint256 payment)
    internal
    returns (bytes32 requestId)
  {
    requestId = keccak256(abi.encodePacked(this, requestCount));
    req.nonce = requestCount;
    pendingRequests[requestId] = oracleAddress;
    emit SeerlinkRequested(requestId);
    require(link.transferAndCall(oracleAddress, payment, encodeRequest(req, ARGS_VERSION)), "unable to transferAndCall to oracle");
    requestCount += 1;

    return requestId;
  }

  /**
   * @notice Allows a request to be cancelled if it has not been fulfilled
   * @dev Requires keeping track of the expiration value emitted from the oracle contract.
   * Deletes the request from the `pendingRequests` mapping.
   * Emits SeerlinkCancelled event.
   * @param requestId The request ID
   * @param payment The amount of LINK sent for the request
   * @param callbackFunc The callback function specified for the request
   * @param expiration The time of the expiration for the request
   */
  function cancelSeerlinkRequest(
    bytes32 requestId,
    uint256 payment,
    bytes4 callbackFunc,
    uint256 expiration
  )
    internal
  {
    SeerlinkRequestInterface requested = SeerlinkRequestInterface(pendingRequests[requestId]);
    delete pendingRequests[requestId];
    emit SeerlinkCancelled(requestId);
    requested.cancelOracleRequest(requestId, payment, callbackFunc, expiration);
  }

  /**
   * @notice Sets the stored oracle address
   * @param oracleAddress The address of the oracle contract
   */
  function setSeerlinkOracle(address oracleAddress) internal {
    oracle = SeerlinkRequestInterface(oracleAddress);
  }

  /**
   * @notice Sets the LINK token address
   * @param linkAddress The address of the LINK token contract
   */
  function setSeerlinkToken(address linkAddress) internal {
    link = LinkTokenInterface(linkAddress);
  }

  /**
   * @notice Sets the Seerlink token address for the public
   * network as given by the Pointer contract
   */
  function setPublicSeerlinkToken() internal {
    setSeerlinkToken(PointerInterface(LINK_TOKEN_POINTER).getAddress());
  }

  /**
   * @notice Retrieves the stored address of the LINK token
   * @return The address of the LINK token
   */
  function seerlinkTokenAddress()
    internal
    view
    returns (address)
  {
    return address(link);
  }

  /**
   * @notice Retrieves the stored address of the oracle contract
   * @return The address of the oracle contract
   */
  function seerlinkOracleAddress()
    internal
    view
    returns (address)
  {
    return address(oracle);
  }

  /**
   * @notice Allows for a request which was created on another contract to be fulfilled
   * on this contract
   * @param oracleAddress The address of the oracle contract that will fulfill the request
   * @param requestId The request ID used for the response
   */
  function addSeerlinkExternalRequest(address oracleAddress, bytes32 requestId)
    internal
    notPendingRequest(requestId)
  {
    pendingRequests[requestId] = oracleAddress;
  }

  /**
   * @notice Sets the stored oracle and LINK token contracts with the addresses resolved by ENS
   * @dev Accounts for subnodes having different resolvers
   * @param ensAddress The address of the ENS contract
   * @param node The ENS node hash
   */
  function useSeerlinkWithENS(address ensAddress, bytes32 node)
    internal
  {
    ens = ENSInterface(ensAddress);
    ensNode = node;
    bytes32 linkSubnode = keccak256(abi.encodePacked(ensNode, ENS_TOKEN_SUBNAME));
    ENSResolver_Seerlink resolver = ENSResolver_Seerlink(ens.resolver(linkSubnode));
    setSeerlinkToken(resolver.addr(linkSubnode));
    updateSeerlinkOracleWithENS();
  }

  /**
   * @notice Sets the stored oracle contract with the address resolved by ENS
   * @dev This may be called on its own as long as `useSeerlinkWithENS` has been called previously
   */
  function updateSeerlinkOracleWithENS()
    internal
  {
    bytes32 oracleSubnode = keccak256(abi.encodePacked(ensNode, ENS_ORACLE_SUBNAME));
    ENSResolver_Seerlink resolver = ENSResolver_Seerlink(ens.resolver(oracleSubnode));
    setSeerlinkOracle(resolver.addr(oracleSubnode));
  }

  /**
   * @notice Encodes the request to be sent to the oracle contract
   * @dev The Seerlink node expects values to be in order for the request to be picked up. Order of types
   * will be validated in the oracle contract.
   * @param req The initialized Seerlink Request
   * @param dataVersion The request data version
   * @return The bytes payload for the `transferAndCall` method
   */
  function encodeRequest(Seerlink.Request memory req, uint256 dataVersion)
    private
    view
    returns (bytes memory)
  {
    return abi.encodeWithSelector(
      oracle.oracleRequest.selector,
      SENDER_OVERRIDE, // Sender value - overridden by onTokenTransfer by the requesting contract's address
      AMOUNT_OVERRIDE, // Amount value - overridden by onTokenTransfer by the actual amount of LINK sent
      req.id,
      req.callbackAddress,
      req.callbackFunctionId,
      req.nonce,
      dataVersion,
      req.buf.buf);
  }

  /**
   * @notice Ensures that the fulfillment is valid for this contract
   * @dev Use if the contract developer prefers methods instead of modifiers for validation
   * @param requestId The request ID for fulfillment
   */
  function validateSeerlinkCallback(bytes32 requestId)
    internal
    recordSeerlinkFulfillment(requestId)
    // solhint-disable-next-line no-empty-blocks
  {}

  /**
   * @dev Reverts if the sender is not the oracle of the request.
   * Emits SeerlinkFulfilled event.
   * @param requestId The request ID for fulfillment
   */
  modifier recordSeerlinkFulfillment(bytes32 requestId) {
    require(msg.sender == pendingRequests[requestId],
            "Source must be the oracle of the request");
    delete pendingRequests[requestId];
    emit SeerlinkFulfilled(requestId);
    _;
  }

  /**
   * @dev Reverts if the request is already pending
   * @param requestId The request ID for fulfillment
   */
  modifier notPendingRequest(bytes32 requestId) {
    require(pendingRequests[requestId] == address(0), "Request is already pending");
    _;
  }
}
