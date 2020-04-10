pragma solidity >=0.4.21 <0.7.0;

import "./util/SafeMath.sol";
import "./SignatureVerifier.sol";
import "./util/Owned.sol";

/**
 * Description: The contract of Sortition.
 **/
contract Sortition is SignatureVerifier{

    using SafeMath256 for uint;
    using SafeMath32 for uint32;
    using SafeMath8 for uint8;

    // ==========
    // STRUCTURE AND STATE VARIABLE:
    // ==========

    uint public randomMod = 100; //Default randon mod data is 100.
    uint public intervalBlock = 100;
    uint public createBlock;

    struct Participator {
        address account;    // participator's account in ethereum.
        uint    factor;     // factor will be used to generate random number.
        bytes32   hash;      // hash of the factor, this param is used for proving
                            // the true factor after submitting factor;
        bytes   signature;  // the signature of account
        bool reveal;
    }
    
    Participator[] participators;

    function addParticapator(
        address _account
    )
        internal
    {
        Participator memory p = Participator({
            account: _account,
            factor: 0,
            hash: "",
            signature: "",
            reveal: false
        });
        participators.push(p);
    }

    /*** MODIFIER ***/
    modifier RandomNumberIsReady(){
        bool allIsReveal = true;
        for(uint i = 0 ; i < participators.length ;i++){
            if( !participators[i].reveal) allIsReveal = false;
        }

        if(allIsReveal){
            _;
        }else{
            require(block.number >= createBlock + intervalBlock);
            _;
        }
    }


    /**
     * description: Upload a factor's hash to the sortition.
     * params:
     * return:
     **/
    function uploadHash(
        bytes32 _hash, 
        bytes _signature, 
        uint _userId)
        external
    {
        if(verifySignature(_hash, _signature, participators[_userId].account)){
            participators[_userId].hash = _hash;
            participators[_userId].signature = _signature;
        }
    }
    
    function uploadFactor(uint _factor, uint _userId)
        external
        returns(bool)
    {
        bytes32 _hash = keccak256(abi.encodePacked(_factor));
        if(_hash == participators[_userId].hash) {
            participators[_userId].factor = _factor.mod(randomMod);
            participators[_userId].reveal = true;
            return true;
        }
        return false;
    }

    function getRandomNum()
        public
        view
        RandomNumberIsReady
        returns(uint)
    {
        uint factorTemp = 0;
        for(uint i = 0; i < participators.length; i++)
        {
            factorTemp = factorTemp.add(participators[i].factor);
            factorTemp = factorTemp.mod(randomMod);
        }
        factorTemp = factorTemp.add(randomMod);
        return factorTemp;
    }
}