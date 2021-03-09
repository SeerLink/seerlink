pragma solidity 0.4.24;

import "./Consumer.sol";

contract UpdatableConsumer is Consumer {

  constructor(bytes32 _specId, address _ens, bytes32 _node) public {
    specId = _specId;
    useSeerlinkWithENS(_ens, _node);
  }

  function updateOracle() public {
    updateSeerlinkOracleWithENS();
  }

  function getSeerlinkToken() public view returns (address) {
    return seerlinkTokenAddress();
  }

  function getOracle() public view returns (address) {
    return seerlinkOracleAddress();
  }

}
