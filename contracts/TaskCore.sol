pragma solidity >=0.4.21 <0.7.0;

import "./User.sol";
import "./QVote.sol";
import "./TaskBase.sol";

contract TaskCore is TaskBase, User {

    uint constant feeReward = 10;
    
    /// EVENT ///
    event CreateTask(address creator,uint taskId);

    /// @dev Create a task
    function createTask(
        string  _taskName,
        string  _content,
        string  _detailContentHash,
        uint _workerLevel
    )
        external
        whenNotPaused
        returns(uint)
    {
        uint taskId = _createTask(_taskName, _content, _detailContentHash, _workerLevel);
        emit CreateTask(msg.sender, taskId);
        return taskId;
    }


    /**
     * description: Accept task by a worker 
     * params: 
     * return: 
     **/
    function acceptTask(uint _taskId)
    external
    returns(bool)
    {
        uint level = getWorkerLevel(msg.sender);
        if(level >= tasks[_taskId].workerLevel){
            //task base 
            _acceptTask( _taskId);
            //user base
            acceptWork(msg.sender);
            return true;
        } else {
            return false;
        }
    }
     

    function submitTaskResult(
        uint _taskId,
        string  _resHash
        )
        external
    {
        _submitTaskResult(_taskId, _resHash);
    }

    /**
     * description: Settlement stage. Task proposer can
     conduct task settlement
     * params:
     * return:
     **/
     function settlementTask(
         uint _taskId,
         uint _score
     )
         payable
         external
     {
         //taskBase
         address _to;
         uint _wei;
        (_to, _wei) =  _settlementTask(_taskId);
        _to.transfer(_wei);
         //userBase
         setUserTaskScore(_to, _score);
     }

    // ==========
    // VOTE INTERFACE:
    // ==========

    /**
     * description: If the requester is not satisfied,
      a deposit is submitted for voting.
     * params:
     * return:
     **/
     function proposeToVote(
        uint _taskId,
        string _requesterProof
     )
        payable
        external
     {
         address _r;
         address _w; 
         (_r,_w)= getVoteFactor(_taskId);
         QVote vote = new QVote(_r,_w);
         _proposeToVote(_taskId, vote, _requesterProof);
     }
    
    /**
     * description: Workers submit deposits for vote against requester
     * params: 
     * return: 
     **/
     function workerAddDeposit(
        uint _taskId,
        string _workerProof
     )
        payable
        public
     {
         QVote vote = QVote(_workerAddDeposit(_taskId, _workerProof));
         vote.voteCommitPhase();
     }

    /// @dev voteForTask (taskid,numvotes)
    // 1. get user msg from user base.( register time, honor)
    // 2. if and only if user who satisfied vote.commitVoteAux could vote
    // 3. deduct the fee first and then deduct msg.value, return if value is more
    // 4. commit vote (commitment , numvote) ,the commitment == keccak256(abi.encodePacked(uint(numVotes),_vote,_salt)
    // _vote is who you support , _salt is a secrte num only you know, in case of someone  guess the result
    function voteForTask(uint _taskId, bytes32 _commitment,uint _numvotes)
        public
        payable
    returns(bool)
    {
        uint registerTime;
        uint honor;
        uint userFee;
        (registerTime,honor,userFee) = getUserVoteFactor(msg.sender);
        
        bool toVote;
        uint voteCost;
        QVote vote = QVote(getTaskVoteAddress(_taskId));
        (toVote,voteCost) = vote.commitVoteAux(msg.sender, registerTime, honor, _numvotes);
        // If user does not satisfied vote sortition, return eth.
        if(!toVote ||(userFee + msg.value < voteCost))
        {
            msg.sender.transfer(msg.value);
            return false;
        }   
        
        // If user's voteFee is enough, return eth;
        // In the vote phase, the fee will transfer from userbase to taskbase
        if(userFee >= voteCost) 
        {
            msg.sender.transfer(msg.value);
            deductUserFee(msg.sender, voteCost); // userbase
            addUserFee(_taskId, voteCost); //taskbase
        }else{
            uint weiForVote = voteCost - userFee;
            msg.sender.transfer(msg.value - weiForVote);
            deductUserFee(msg.sender, userFee); // userbase
            addUserFee(_taskId, userFee); //taskbase
            addUserWei(_taskId, weiForVote); //taskbase
        }
        
        vote.commitVote(msg.sender, _commitment, uint32(_numvotes));
    }   
    
    function revealVoteForTask(uint _taskId, uint _vote, bytes32 _salt)
        public
    {
        QVote vote = QVote(getTaskVoteAddress(_taskId));
        vote.revealVote(msg.sender, _vote, _salt);
        // add eth wei to the special side
        _transferWeiToSupport(_taskId, _vote);
    }
    
    
    
    /// @dev In the end of vote phase, the vote fee transfers from taskbase to userbase
    function withdrawVoteDeposit(uint _taskId)
        private
    {
        uint _fee = subUserFee(_taskId); //taskbase
        increaseUserFee(msg.sender, _fee); //userbase
        uint _wei = subUserWei(_taskId); // taskbase
        msg.sender.transfer(_wei); //userbase
    }
    
    
    /// @dev the voter who support the winner could get reward
    function withdrawVoteReward(uint _taskId)
        public
        returns(bool)
    {
        QVote vote = QVote(getTaskVoteAddress(_taskId));
        if(!vote.isVoteCompleted())
            return false;
        bool isVoter; 
        bool isWinSide;  
        uint winChoice;
        (isVoter, isWinSide, winChoice) = vote.getVoterResult(msg.sender);
        if(isVoter && isWinSide) {
            // get task reward
            uint sharedReward = getSharedReward(_taskId, winChoice);
            msg.sender.transfer(sharedReward);
            // get back vote deposit. This function must be after user get shareReward.
            withdrawVoteDeposit(_taskId);
            // user get system reward
            increaseUserFee(msg.sender, feeReward);
            return true;
        }
        return false;
    }
    
    
    /// @dev task requester or worker who win the poll could get task reward
    function withdrawTaskReward(uint _taskId)
        public
        returns(bool)
    {
        QVote vote = QVote(getTaskVoteAddress(_taskId));
        if(!vote.isVoteCompleted())
            return false;   
        uint winChoice = vote.completeVote();
        (address _to, uint _wei) = getTaskReward(_taskId, winChoice);
        _to.transfer(_wei);
        return true;
    }
    
}

