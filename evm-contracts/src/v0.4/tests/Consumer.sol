pragma solidity 0.4.24;

import "../SeerlinkClient.sol";

contract Consumer is SeerlinkClient {
  bytes32 internal specId;
  bytes32 public currentPrice;

  uint256 constant private ORACLE_PAYMENT = 1 * LINK;

  event RequestFulfilled(
    bytes32 indexed requestId,  // User-defined ID
    bytes32 indexed price
  );

  function requestEthereumPrice(string _currency) public {
    Seerlink.Request memory req = buildSeerlinkRequest(specId, this, this.fulfill.selector);
    req.add("get", "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD,EUR,JPY");
    string[] memory path = new string[](1);
    path[0] = _currency;
    req.addStringArray("path", path);
    sendSeerlinkRequest(req, ORACLE_PAYMENT);
  }

  function cancelRequest(
    bytes32 _requestId,
    uint256 _payment,
    bytes4 _callbackFunctionId,
    uint256 _expiration
  ) public {
    cancelSeerlinkRequest(_requestId, _payment, _callbackFunctionId, _expiration);
  }

  function withdrawLink() public {
    LinkTokenInterface link = LinkTokenInterface(seerlinkTokenAddress());
    require(link.transfer(msg.sender, link.balanceOf(address(this))), "Unable to transfer");
  }

  function fulfill(bytes32 _requestId, bytes32 _price)
    public
    recordSeerlinkFulfillment(_requestId)
  {
    emit RequestFulfilled(_requestId, _price);
    currentPrice = _price;
  }

}
