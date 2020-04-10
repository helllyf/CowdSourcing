pragma solidity >=0.4.21 <0.7.0;

/// @dev A facet of CrowdSourcing that manages special access privileges.
contract AccessControl {

    /// @dev Emitted when contract is upgraded.
    event ContractUpgrade(address newContract);

    /*** STATE VARIABLE ***/
    /// @dev The address of account that can execute actions.
    address public ceoAddress;

    /// @dev Keeps track whether the contract is paused. When that is true, most of actions are blocks.
    bool public paused = false;


    /*** MODIFIER ***/
    
    constructor() public{
        ceoAddress = msg.sender;
    }
    
    /// @dev Access modifier for CEO-only functionality
    modifier onlyCEO() {
        require(msg.sender == ceoAddress);
        _;
    }
    
    /// @dev Modifier to allow actions only when the contract IS NOT paused.
    modifier whenNotPaused() {
        require(!paused);
        _;
    }

    /// @dev Modifier to allow actions only when the contract IS paused.
    modifier whenPaused() {
        require(paused);
        _;
    }
    
    /// @dev Assigns a new address to act as CEO. Only avaliable to the current CEO.
    /// @param _newCEO The new address of CEO
    function setCEO(address _newCEO)
        external
        onlyCEO
    {
        require(_newCEO != address(0));
        ceoAddress = _newCEO;
    }

    /// @dev Call by CEO role to pause the contract. Only when a bug or exploit is detected and we need to limit damage.
    function pause()
        external
        onlyCEO
        whenNotPaused
    {
        paused = true;
    }

    /// @dev Call by CEO role to unpause the contract. 
    /// @notice This is public rather than external so it can be call by derived contracts.
    function unpause()
        public
        onlyCEO
        whenPaused
    {
        paused = false;
    }
}

