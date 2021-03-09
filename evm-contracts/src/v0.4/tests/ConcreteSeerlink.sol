pragma solidity 0.4.24;

import "../Seerlink.sol";
import { CBOR as CBOR_Seerlink } from "../vendor/CBOR.sol";
import { Buffer as Buffer_Seerlink } from "../vendor/Buffer.sol";

contract ConcreteSeerlink {
  using Seerlink for Seerlink.Request;
  using CBOR_Seerlink for Buffer_Seerlink.buffer;

  Seerlink.Request private req;

  event RequestData(bytes payload);

  function closeEvent() public {
    emit RequestData(req.buf.buf);
  }

  function setBuffer(bytes data) public {
    Seerlink.Request memory r2 = req;
    r2.setBuffer(data);
    req = r2;
  }

  function add(string _key, string _value) public {
    Seerlink.Request memory r2 = req;
    r2.add(_key, _value);
    req = r2;
  }

  function addBytes(string _key, bytes _value) public {
    Seerlink.Request memory r2 = req;
    r2.addBytes(_key, _value);
    req = r2;
  }

  function addInt(string _key, int256 _value) public {
    Seerlink.Request memory r2 = req;
    r2.addInt(_key, _value);
    req = r2;
  }

  function addUint(string _key, uint256 _value) public {
    Seerlink.Request memory r2 = req;
    r2.addUint(_key, _value);
    req = r2;
  }

  // Temporarily have method receive bytes32[] memory until experimental
  // string[] memory can be invoked from truffle tests.
  function addStringArray(string _key, bytes32[] memory _values) public {
    string[] memory strings = new string[](_values.length);
    for (uint256 i = 0; i < _values.length; i++) {
      strings[i] = bytes32ToString(_values[i]);
    }
    Seerlink.Request memory r2 = req;
    r2.addStringArray(_key, strings);
    req = r2;
  }

  function bytes32ToString(bytes32 x) private pure returns (string) {
    bytes memory bytesString = new bytes(32);
    uint charCount = 0;
    for (uint j = 0; j < 32; j++) {
      byte char = byte(bytes32(uint(x) * 2 ** (8 * j)));
      if (char != 0) {
        bytesString[charCount] = char;
        charCount++;
      }
    }
    bytes memory bytesStringTrimmed = new bytes(charCount);
    for (j = 0; j < charCount; j++) {
      bytesStringTrimmed[j] = bytesString[j];
    }
    return string(bytesStringTrimmed);
  }
}
