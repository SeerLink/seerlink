pragma solidity ^0.4.24;

import "./SeerlinkClient.sol";

/**
 * @title The Seerlinked contract
 * @notice Contract writers can inherit this contract in order to create requests for the
 * Seerlink network. SeerlinkClient is an alias of the Seerlinked contract.
 */
contract Seerlinked is SeerlinkClient {
  /**
   * @notice Creates a request that can hold additional parameters
   * @param _specId The Job Specification ID that the request will be created for
   * @param _callbackAddress The callback address that the response will be sent to
   * @param _callbackFunctionSignature The callback function signature to use for the callback address
   * @return A Seerlink Request struct in memory
   */
  function newRequest(
    bytes32 _specId,
    address _callbackAddress,
    bytes4 _callbackFunctionSignature
  ) internal pure returns (Seerlink.Request memory) {
    return buildSeerlinkRequest(_specId, _callbackAddress, _callbackFunctionSignature);
  }

  /**
   * @notice Creates a Seerlink request to the stored oracle address
   * @dev Calls `sendSeerlinkRequestTo` with the stored oracle address
   * @param _req The initialized Seerlink Request
   * @param _payment The amount of LINK to send for the request
   * @return The request ID
   */
  function seerlinkRequest(Seerlink.Request memory _req, uint256 _payment)
    internal
    returns (bytes32)
  {
    return sendSeerlinkRequest(_req, _payment);
  }

  /**
   * @notice Creates a Seerlink request to the specified oracle address
   * @dev Generates and stores a request ID, increments the local nonce, and uses `transferAndCall` to
   * send LINK which creates a request on the target oracle contract.
   * Emits SeerlinkRequested event.
   * @param _oracle The address of the oracle for the request
   * @param _req The initialized Seerlink Request
   * @param _payment The amount of LINK to send for the request
   * @return The request ID
   */
  function seerlinkRequestTo(address _oracle, Seerlink.Request memory _req, uint256 _payment)
    internal
    returns (bytes32 requestId)
  {
    return sendSeerlinkRequestTo(_oracle, _req, _payment);
  }

  /**
   * @notice Sets the stored oracle address
   * @param _oracle The address of the oracle contract
   */
  function setOracle(address _oracle) internal {
    setSeerlinkOracle(_oracle);
  }

  /**
   * @notice Sets the LINK token address
   * @param _link The address of the LINK token contract
   */
  function setLinkToken(address _link) internal {
    setSeerlinkToken(_link);
  }

  /**
   * @notice Retrieves the stored address of the LINK token
   * @return The address of the LINK token
   */
  function seerlinkToken()
    internal
    view
    returns (address)
  {
    return seerlinkTokenAddress();
  }

  /**
   * @notice Retrieves the stored address of the oracle contract
   * @return The address of the oracle contract
   */
  function oracleAddress()
    internal
    view
    returns (address)
  {
    return seerlinkOracleAddress();
  }

  /**
   * @notice Ensures that the fulfillment is valid for this contract
   * @dev Use if the contract developer prefers methods instead of modifiers for validation
   * @param _requestId The request ID for fulfillment
   */
  function fulfillSeerlinkRequest(bytes32 _requestId)
    internal
    recordSeerlinkFulfillment(_requestId)
    // solhint-disable-next-line no-empty-blocks
  {}

  /**
   * @notice Sets the stored oracle and LINK token contracts with the addresses resolved by ENS
   * @dev Accounts for subnodes having different resolvers
   * @param _ens The address of the ENS contract
   * @param _node The ENS node hash
   */
  function setSeerlinkWithENS(address _ens, bytes32 _node)
    internal
  {
    useSeerlinkWithENS(_ens, _node);
  }

  /**
   * @notice Sets the stored oracle contract with the address resolved by ENS
   * @dev This may be called on its own as long as `setSeerlinkWithENS` has been called previously
   */
  function setOracleWithENS()
    internal
  {
    updateSeerlinkOracleWithENS();
  }

  /**
   * @notice Allows for a request which was created on another contract to be fulfilled
   * on this contract
   * @param _oracle The address of the oracle contract that will fulfill the request
   * @param _requestId The request ID used for the response
   */
  function addExternalRequest(address _oracle, bytes32 _requestId)
    internal
  {
    addSeerlinkExternalRequest(_oracle, _requestId);
  }
}
