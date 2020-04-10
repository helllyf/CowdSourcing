pragma solidity >=0.4.21 <0.7.0;

import "./util/SafeMath.sol";



contract UserBase {
    using SafeMath256 for uint;
    using SafeMath32 for uint32;
    using SafeMath8 for uint8;

    /*** EVENT ***/
    event RegisterUser(address user);

    event AcceptWork(address user, uint acceptwork);

    event FinishWork(address user, uint finishedwork, uint workscore);
    
    /*** DATA STRUCTURE AND STATE VARIBALE ***/

    struct UserTuple {
        /*** Task Data ***/
        uint registerTime;
        uint acceptedWork;
        uint workScore;     // Average Score = workScore / finishTask ,which is work level
        uint finishedWork;   //Task finsih rate = finishTask / acceptTask

        /*** Voter Data ***/
        uint voteWin;  // The number of a user participates as a voter and wins
        uint voteRewardFee; // Users are rewarded with token for voting, which can be used as an alternative to ETH.
        //uint voteRewardWei; // Users can withdraw the rewards they receive for participating in the vote.

        /*** Vote Participator Data ***/
        uint candidateAsWorker;
        uint candidateAsRequester;
        
    }

    mapping(address => UserTuple) public users;

    function _register(address _user)
        internal
    {
        // the user need unregister
        require(users[_user].registerTime == 0);
        
        UserTuple memory _userTuple = UserTuple({
            registerTime: now,
            workScore:  0,
            acceptedWork: 0,
            finishedWork: 0,
            voteWin: 0,
            voteRewardFee: 0,
            //voteRewardWei: 0,
            candidateAsWorker: 0,
            candidateAsRequester: 0
        });

        users[_user] = _userTuple;

        emit RegisterUser(_user);
    }

    function _acceptWork(address _user)
        internal
    {
        users[_user].acceptedWork = users[_user].acceptedWork.add(1);
        emit AcceptWork(_user, users[_user].acceptedWork);
    }

    function _finishWork(address _user, uint _workScore)
        internal
    {
        users[_user].finishedWork = users[_user].finishedWork.add(1);
        users[_user].workScore = users[_user].workScore.add(_workScore);

        emit FinishWork(_user, users[_user].finishedWork, users[_user].workScore);
    }

    function _userAsCandidate(address _requester, address _worker)
        internal
    {
        users[_worker].candidateAsWorker = users[_worker].candidateAsWorker.add(1);
        users[_requester].candidateAsRequester = users[_requester].candidateAsRequester.add(1);
    }

    /// @dev the facet of user fee
    function _increaseUserFee(address _user, uint _fee)
        internal
    {
        users[_user].voteRewardFee = users[_user].voteRewardFee.add(_fee); 
    }
    
    /// @dev the facet of user fee
    function _deductUserFee(address _user, uint _fee)
        internal
    {
        users[_user].voteRewardFee = users[_user].voteRewardFee.sub(_fee); 
    }
}


contract User is UserBase{
    using SafeMath256 for uint;
    using SafeMath32 for uint32;
    using SafeMath8 for uint8;

    // ==========
    // DATA STRUCTURE AND STATE VARIBALE:
    // ==========

    // ==========
    // USER INTERFACE:
    // ==========

    function registerUser()
        public
    {
         _register(msg.sender);
    }

    /**
     * description: Get worker level
     * params:
     * return:
     **/
     function getWorkerLevel(address _worker)
        public
        view
        returns(uint)
     {
        uint l;
        if( 0 == users[_worker].finishedWork )
            l = 0;
        else
            l = users[_worker].workScore.div(users[_worker].finishedWork);
        return l;
     }

    function setUserTaskScore(address _worker, uint _score)
        internal
    {
        _finishWork(_worker,_score);
    }
    
    function acceptWork(address _user)  
        internal
    {
        _acceptWork(_user);
    }
    
    function getUserVoteFactor(address _user)
        view
        internal
    returns(uint _registerTime, uint _honor, uint _userFee)
    {
        _registerTime = users[_user].registerTime;
        _userFee = users[_user].voteRewardFee;
        /// @dev honor is between 0 ~ 40
        
        // task Score, the task score up limit is 10
        // each vote, 10 fee  will be reward
        uint _userScoreHonor = 20 * (users[_user].workScore / users[_user].finishedWork) / 10;
        uint _userVoteHonor = 20 * (users[_user].voteRewardFee) / 200;
        _honor = _userScoreHonor + _userVoteHonor;
    }
    
 
    function deductUserFee(address _user, uint _fee)
        internal
    {
         _deductUserFee(_user, _fee);
    }
    
    function increaseUserFee(address _user, uint _fee)
        internal
    {
        _increaseUserFee(_user, _fee);
    }
}

