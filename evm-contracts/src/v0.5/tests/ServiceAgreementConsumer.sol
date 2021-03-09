pragma solidity 0.5.0;

import "../SeerlinkClient.sol";

contract ServiceAgreementConsumer is SeerlinkClient {
  uint256 constant private ORACLE_PAYMENT = 1 * LINK;

  bytes32 internal sAId;
  bytes32 public currentPrice;

  constructor(address _link, address _coordinator, bytes32 _sAId) public {
    setSeerlinkToken(_link);
    setSeerlinkOracle(_coordinator);
    sAId = _sAId;
  }

  function requestEthereumPrice(string memory _currency) public {
    Seerlink.Request memory req = buildSeerlinkRequest(sAId, address(this), this.fulfill.selector);
    req.add("get", "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD,EUR,JPY");
    req.add("path", _currency);
    sendSeerlinkRequest(req, ORACLE_PAYMENT);
  }

  function fulfill(bytes32 _requestId, bytes32 _price)
    public
    recordSeerlinkFulfillment(_requestId)
  {
    currentPrice = _price;
  }
}
