// Specifies the version of Solidity, using semantic versioning.
// Learn more: https://solidity.readthedocs.io/en/v0.5.10/layout-of-source-files.html#pragma
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

// Defines a contract named `HelloWorld`.
// A contract is a collection of functions and data (its state). Once deployed, a contract resides at a specific address on the Ethereum blockchain. Learn more: https://solidity.readthedocs.io/en/v0.5.10/structure-of-a-contract.html
contract HelloWorld is Ownable {
    // Hello function available for everyone.
    function Hello() public pure returns (string memory) {
        return "Hello World";
    }

    // GreetOwner function is only available for the owner.
    function GreetOwner(string memory name)
        public
        view
        onlyOwner
        returns (string memory)
    {
        return string(abi.encodePacked("Hello ", name, "!"));
    }
}
