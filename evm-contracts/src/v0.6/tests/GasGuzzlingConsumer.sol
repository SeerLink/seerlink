// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./Consumer.sol";

contract GasGuzzlingConsumer is Consumer{

  constructor(address _link, address _oracle, bytes32 _specId) public {
    setSeerlinkToken(_link);
    setSeerlinkOracle(_oracle);
    specId = _specId;
  }

  function gassyRequestEthereumPrice(uint256 _payment) public {
    Seerlink.Request memory req = buildSeerlinkRequest(specId, address(this), this.gassyFulfill.selector);
    req.add("get", "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD,EUR,JPY");
    string[] memory path = new string[](1);
    path[0] = "USD";
    req.addStringArray("path", path);
    sendSeerlinkRequest(req, _payment);
  }

  function gassyFulfill(bytes32 _requestId, bytes32 _price)
    public
    recordSeerlinkFulfillment(_requestId)
  {
    while(true){
    }
  }

  function gassyMultiWordRequest(uint256 _payment) public {
    Seerlink.Request memory req = buildSeerlinkRequest(specId, address(this), this.gassyMultiWordFulfill.selector);
    req.add("get", "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD,EUR,JPY");
    string[] memory path = new string[](1);
    path[0] = "USD";
    req.addStringArray("path", path);
    sendSeerlinkRequest(req, _payment);
  }

  function gassyMultiWordFulfill(bytes32 _requestId, bytes memory _price)
    public
    recordSeerlinkFulfillment(_requestId)
  {
    while(true){
    }
  }
}