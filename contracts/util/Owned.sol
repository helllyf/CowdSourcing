pragma solidity >=0.4.21 <0.7.0;

/**
@title Owned
@dev establishes contract owner, implements modifier to require owner 
status, and allows transfer of ownership
*/
contract Owned {

    /// @dev This is the deployer of the contract, which may be the user or the contract address.
    address public owner;

    constructor() internal { owner = msg.sender; }
    modifier onlyOwner { require (msg.sender == owner); _; }

    function transferOwnership(address newOwner) onlyOwner public { 
        owner = newOwner; 
    }
}