pragma solidity 0.5.0;

import "./MaliciousSeerlink.sol";
import "../SeerlinkClient.sol";
import "../vendor/SafeMathSeerlink.sol";

contract MaliciousSeerlinkClient is SeerlinkClient {
  using MaliciousSeerlink for MaliciousSeerlink.Request;
  using MaliciousSeerlink for MaliciousSeerlink.WithdrawRequest;
  using Seerlink for Seerlink.Request;
  using SafeMathSeerlink for uint256;

  uint256 private maliciousRequests = 1;
  mapping(bytes32 => address) private maliciousPendingRequests;

  function newWithdrawRequest(
    bytes32 _specId,
    address _callbackAddress,
    bytes4 _callbackFunction
  ) internal pure returns (MaliciousSeerlink.WithdrawRequest memory) {
    MaliciousSeerlink.WithdrawRequest memory req;
    return req.initializeWithdraw(_specId, _callbackAddress, _callbackFunction);
  }

  function seerlinkTargetRequest(address _target, Seerlink.Request memory _req, uint256 _amount)
    internal
    returns(bytes32 requestId)
  {
    requestId = keccak256(abi.encodePacked(_target, maliciousRequests));
    _req.nonce = maliciousRequests;
    maliciousPendingRequests[requestId] = seerlinkOracleAddress();
    emit SeerlinkRequested(requestId);
    LinkTokenInterface _link = LinkTokenInterface(seerlinkTokenAddress());
    require(_link.transferAndCall(seerlinkOracleAddress(), _amount, encodeTargetRequest(_req)), "Unable to transferAndCall to oracle");
    maliciousRequests += 1;

    return requestId;
  }

  function seerlinkPriceRequest(Seerlink.Request memory _req, uint256 _amount)
    internal
    returns(bytes32 requestId)
  {
    requestId = keccak256(abi.encodePacked(this, maliciousRequests));
    _req.nonce = maliciousRequests;
    maliciousPendingRequests[requestId] = seerlinkOracleAddress();
    emit SeerlinkRequested(requestId);
    LinkTokenInterface _link = LinkTokenInterface(seerlinkTokenAddress());
    require(_link.transferAndCall(seerlinkOracleAddress(), _amount, encodePriceRequest(_req)), "Unable to transferAndCall to oracle");
    maliciousRequests += 1;

    return requestId;
  }

  function seerlinkWithdrawRequest(MaliciousSeerlink.WithdrawRequest memory _req, uint256 _wei)
    internal
    returns(bytes32 requestId)
  {
    requestId = keccak256(abi.encodePacked(this, maliciousRequests));
    _req.nonce = maliciousRequests;
    maliciousPendingRequests[requestId] = seerlinkOracleAddress();
    emit SeerlinkRequested(requestId);
    LinkTokenInterface _link = LinkTokenInterface(seerlinkTokenAddress());
    require(_link.transferAndCall(seerlinkOracleAddress(), _wei, encodeWithdrawRequest(_req)), "Unable to transferAndCall to oracle");
    maliciousRequests += 1;
    return requestId;
  }

  function encodeWithdrawRequest(MaliciousSeerlink.WithdrawRequest memory _req)
    internal pure returns (bytes memory)
  {
    return abi.encodeWithSelector(
      bytes4(keccak256("withdraw(address,uint256)")),
      _req.callbackAddress,
      _req.callbackFunctionId,
      _req.nonce,
      _req.buf.buf);
  }

  function encodeTargetRequest(Seerlink.Request memory _req)
    internal pure returns (bytes memory)
  {
    return abi.encodeWithSelector(
      bytes4(keccak256("oracleRequest(address,uint256,bytes32,address,bytes4,uint256,uint256,bytes)")),
      0, // overridden by onTokenTransfer
      0, // overridden by onTokenTransfer
      _req.id,
      _req.callbackAddress,
      _req.callbackFunctionId,
      _req.nonce,
      1,
      _req.buf.buf);
  }

  function encodePriceRequest(Seerlink.Request memory _req)
    internal pure returns (bytes memory)
  {
    return abi.encodeWithSelector(
      bytes4(keccak256("oracleRequest(address,uint256,bytes32,address,bytes4,uint256,uint256,bytes)")),
      0, // overridden by onTokenTransfer
      2000000000000000000, // overridden by onTokenTransfer
      _req.id,
      _req.callbackAddress,
      _req.callbackFunctionId,
      _req.nonce,
      1,
      _req.buf.buf);
  }
}
