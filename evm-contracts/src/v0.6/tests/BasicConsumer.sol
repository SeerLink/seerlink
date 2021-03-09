// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./Consumer.sol";

contract BasicConsumer is Consumer {

  constructor(address _link, address _oracle, bytes32 _specId) public {
    setSeerlinkToken(_link);
    setSeerlinkOracle(_oracle);
    specId = _specId;
  }

}
