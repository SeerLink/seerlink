// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "../SeerlinkClient.sol";
import "../vendor/SafeMathSeerlink.sol";

contract MaliciousMultiWordConsumer is SeerlinkClient {
  using SafeMathSeerlink for uint256;

  uint256 constant private ORACLE_PAYMENT = 1 * LINK;
  uint256 private expiration;

  constructor(address _link, address _oracle) public payable {
    setSeerlinkToken(_link);
    setSeerlinkOracle(_oracle);
  }

  receive() external payable {} // solhint-disable-line no-empty-blocks

  function requestData(bytes32 _id, bytes memory _callbackFunc) public {
    Seerlink.Request memory req = buildSeerlinkRequest(_id, address(this), bytes4(keccak256(_callbackFunc)));
    expiration = now.add(5 minutes); // solhint-disable-line not-rely-on-time
    sendSeerlinkRequest(req, ORACLE_PAYMENT);
  }

  function assertFail(bytes32, bytes memory) public pure {
    assert(1 == 2);
  }

  function cancelRequestOnFulfill(bytes32 _requestId, bytes memory) public {
    cancelSeerlinkRequest(
      _requestId,
      ORACLE_PAYMENT,
      this.cancelRequestOnFulfill.selector,
      expiration);
  }

  function remove() public {
    selfdestruct(address(0));
  }

  function stealEthCall(bytes32 _requestId, bytes memory) public recordSeerlinkFulfillment(_requestId) {
    (bool success,) = address(this).call.value(100)(""); // solhint-disable-line avoid-call-value
    require(success, "Call failed");
  }

  function stealEthSend(bytes32 _requestId, bytes memory) public recordSeerlinkFulfillment(_requestId) {
    // solhint-disable-next-line check-send-result
    bool success = address(this).send(100); // solhint-disable-line multiple-sends
    require(success, "Send failed");
  }

  function stealEthTransfer(bytes32 _requestId, bytes memory) public recordSeerlinkFulfillment(_requestId) {
    address(this).transfer(100);
  }

  function doesNothing(bytes32, bytes memory) public pure {} // solhint-disable-line no-empty-blocks
}
