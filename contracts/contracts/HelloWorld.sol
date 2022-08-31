// Specifies the version of Solidity, using semantic versioning.
// Learn more: https://solidity.readthedocs.io/en/v0.5.10/layout-of-source-files.html#pragma
pragma solidity >=0.7.3;

import "@openzeppelin/contracts/access/Ownable.sol";

// Defines a contract named `HelloWorld`.
// A contract is a collection of functions and data (its state). Once deployed, a contract resides at a specific address on the Ethereum blockchain. Learn more: https://solidity.readthedocs.io/en/v0.5.10/structure-of-a-contract.html
contract HelloWorld {
    function Hello() public pure returns (string memory) {
        return "Hello World";
    }

    function Greet(string memory str)
        public
        pure
        onlyOwner
        returns (string memory)
    {
        return str;
    }
}
