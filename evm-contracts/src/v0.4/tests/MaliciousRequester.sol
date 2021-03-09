pragma solidity 0.4.24;


import "./MaliciousSeerlinked.sol";


contract MaliciousRequester is MaliciousSeerlinked {

  uint256 constant private ORACLE_PAYMENT = 1 * LINK;
  uint256 private expiration;

  constructor(address _link, address _oracle) public {
    setLinkToken(_link);
    setOracle(_oracle);
  }

  function maliciousWithdraw()
    public
  {
    MaliciousSeerlink.WithdrawRequest memory req = newWithdrawRequest(
      "specId", this, this.doesNothing.selector);
    seerlinkWithdrawRequest(req, ORACLE_PAYMENT);
  }

  function request(bytes32 _id, address _target, bytes _callbackFunc) public returns (bytes32 requestId) {
    Seerlink.Request memory req = newRequest(_id, _target, bytes4(keccak256(_callbackFunc)));
    expiration = now.add(5 minutes); // solhint-disable-line not-rely-on-time
    requestId = seerlinkRequest(req, ORACLE_PAYMENT);
  }

  function maliciousPrice(bytes32 _id) public returns (bytes32 requestId) {
    Seerlink.Request memory req = newRequest(_id, this, this.doesNothing.selector);
    requestId = seerlinkPriceRequest(req, ORACLE_PAYMENT);
  }

  function maliciousTargetConsumer(address _target) public returns (bytes32 requestId) {
    Seerlink.Request memory req = newRequest("specId", _target, bytes4(keccak256("fulfill(bytes32,bytes32)")));
    requestId = seerlinkTargetRequest(_target, req, ORACLE_PAYMENT);
  }

  function maliciousRequestCancel(bytes32 _id, bytes _callbackFunc) public {
    SeerlinkRequestInterface oracle = SeerlinkRequestInterface(oracleAddress());
    oracle.cancelOracleRequest(
      request(_id, this, _callbackFunc),
      ORACLE_PAYMENT,
      this.maliciousRequestCancel.selector,
      expiration
    );
  }

  function doesNothing(bytes32, bytes32) public pure {} // solhint-disable-line no-empty-blocks
}
