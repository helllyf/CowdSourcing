pragma solidity >=0.4.21 <0.7.0;

import "./util/SafeMath.sol";
import "./AccessControl.sol";

/// @dev Base contract for Task. Holds all common structs, events and base variables.
contract TaskBase is AccessControl {

    using SafeMath256 for uint;
    using SafeMath32 for uint32;
    using SafeMath8 for uint8;
    
    /*** EVENTS ***/

    /// @dev The CreateTask is fired whenever a new task comes into existence.
    event CreateTask(address requester, uint taskId);

    /// @dev The AcceptTask is fired whenever a task was accepted
    event AcceptTask(address worker, uint taskId);

    /// @dev The SubmitTaskHash is fired whenever the result of task work is finished.
    event SubmitTaskResult(address requester, address worker, uint taskId, string resultHash);

    /// @dev The SettlementTask is fired whenever the task is settlement.
    event SettlementTask(uint taskId);

    /// @dev The ProposedToVote is fired whenever the task requester does not satisfied the work result of worker, and does not want to pay for worker.
    event ProposeToVote(uint taskId, address vote);

    /// @dev The WorkerAddDeposit is fired whenever the worker who want to take against the requester's vote.
    event WorkerAddDeposit(uint taskId, address vote);
    
    /*** DATA TYPE ***/
    enum TaskStatus {createdPhase, commitPhase, settlePhase, votePhase, finishPhase}

    struct Task {
        // The task name, that can't change.
        string taskName;
        // Introduction of the task.
        string content;
        // The detail of the task content, that stores hash of the content in IPFS.
        string detailContentHash;
        //Worker's level need to be satisfied, which is the minimum standard for worker.
        uint workerLevel;
        TaskStatus status;      // The task's phase, which is always changing until the task is finish 
        uint totalWeiPaid;      //Total amount paid into the task (in Wei)
        uint taskRewardWei;     //Paid for the task
        string resultHash;      // Worker submit his work, which is stored in IPFS
        uint256 createTime;      // The task Create time.
        uint256 createBlock;     // The task's create block;
        uint requesterDepositWei;  // Requester's deposit for vote
        uint supportRequesterWei; // wei of voter who support requester
        uint workerDepositWei;    // Worker's deposit for vote
        uint supportWorkerWei;
        address voteAddr;       //This is the vote address in ethereum
        string requesterProof;  //requester's proof for voter, it helps voter to make decision
        string workerProof;     //worker's proof for voter, it helps voter to make decision

        mapping(address => uint) voterFee; //User voting cost, Fee transfer from user base
        mapping(address => uint) voterWei; //User voting cost, Wei is Eth unit
    }

    /*** CONSTANTS ***/

    // An approximation of currently how many seconds are in between blocks.
    uint256 public secondsPerBlock = 15;

    /// @dev The intervalBlock is the interval of each phase.
    uint public intervalBlock = 2;

    /*** STORAGE ***/

    /// @dev An array contains task struct for all tasks in existence. The
    /// ID of each task is actually an index into this array.this
    Task[] tasks;

    /// @dev A mapping from task id to an address that propose this task.
    mapping(uint => address) public taskIndexToRequester;

    /// @dev A mapping from task id to an address that accept this work.
    mapping(uint => address) public taskIdToWorker;

    /// @dev An internal method that create a task and stores it. This
    /// method does not do any checking and should only be called when
    /// the input data is known to be valid.
    /// @param _taskName The name of this task.
    /// @param _content The content of this task.
    /// @param _detailContentHash The hash of the detail content which stores
    /// in IPFS
    /// @param _workerLevel The worker's level need to be satisfied for this task.
    /// returns New created task id will be return by this method.
    function _createTask(
        string memory _taskName,
        string memory _content,
        string memory _detailContentHash,
        uint _workerLevel
    )
    internal
    returns(uint)
    {
        Task memory _task = Task({
            taskName: _taskName,
            content: _content,
            detailContentHash: _detailContentHash,
            workerLevel: _workerLevel,
            status: TaskStatus.createdPhase,
            totalWeiPaid: 0,
            taskRewardWei: 0,
            requesterDepositWei: 0,
            supportRequesterWei: 0,
            workerDepositWei: 0,
            supportWorkerWei: 0,
            createTime: uint64(now),
            createBlock: uint256(block.number),
            voteAddr: address(0),
            requesterProof: "",
            workerProof: "",
            resultHash: ""
            });
        uint256 newTaskId = tasks.push(_task) - 1;
        taskIndexToRequester[newTaskId] = msg.sender;

        emit CreateTask(msg.sender, newTaskId);
    }

    /// @dev An internal method for worker to accept task
    function _acceptTask(
        uint _taskId
    )
        internal
    {
        require(block.number >= tasks[_taskId].createBlock + intervalBlock);
        require(tasks[_taskId].status == TaskStatus.createdPhase);

        taskIdToWorker[_taskId] = msg.sender;
        tasks[_taskId].status = TaskStatus.commitPhase;
        emit AcceptTask(msg.sender, _taskId);
    }

    /// @dev An internal method for worker to submmit work result.
    /// @param _taskId The id of task in array
    /// @param _resultHash The result of working which stores in IPFS
    function _submitTaskResult(
        uint _taskId,
        string memory _resultHash
    )
        internal
    {
        require(tasks[_taskId].status == TaskStatus.commitPhase);
        require(taskIdToWorker[_taskId] == msg.sender);
        tasks[_taskId].resultHash = _resultHash;
        tasks[_taskId].status = TaskStatus.settlePhase;

        emit SubmitTaskResult(
            taskIndexToRequester[_taskId],
            msg.sender,
            _taskId,
            _resultHash
        );
    }


    function _settlementTask(
        uint _taskId
    )
        internal
        returns(address _to, uint _wei)
    {
        require( taskIndexToRequester[_taskId] == msg.sender);
        require( tasks[_taskId].status == TaskStatus.settlePhase );

        //taskIdToWorker[_taskId].transfer(tasks[_taskId].taskRewardWei);
        _to = taskIdToWorker[_taskId];
        _wei = tasks[_taskId].taskRewardWei;
        
        tasks[_taskId].totalWeiPaid = tasks[_taskId].totalWeiPaid.sub(tasks[_taskId].taskRewardWei);
        tasks[_taskId].taskRewardWei = 0;
        tasks[_taskId].status = TaskStatus.finishPhase;

        emit SettlementTask(_taskId);
    }


    function _proposeToVote(
        uint _taskId,
        address _vote,
        string memory _requesterProof
    )
        internal
    {
        require(taskIndexToRequester[_taskId] == msg.sender);
        require(tasks[_taskId].status == TaskStatus.commitPhase);

        tasks[_taskId].status = TaskStatus.votePhase;
        tasks[_taskId].requesterDepositWei = tasks[_taskId].requesterDepositWei.add(msg.value);

        tasks[_taskId].voteAddr = _vote;
        tasks[_taskId].requesterProof = _requesterProof;
        emit ProposeToVote(_taskId, _vote);
    }


    function _workerAddDeposit(
        uint _taskId,
        string memory _workerProof
    )
        internal
        returns(address)
    {
        require(tasks[_taskId].status == TaskStatus.votePhase);
        require(taskIdToWorker[_taskId] == msg.sender);

        tasks[_taskId].workerDepositWei = tasks[_taskId].workerDepositWei.add(msg.value);
        tasks[_taskId].workerProof = _workerProof;

        emit WorkerAddDeposit(_taskId, tasks[_taskId].voteAddr);
        return getTaskVoteAddress(_taskId);
    }

    function _transferWeiToSupport(uint _taskId , uint _vote)
        internal
    {
        if(_vote == 0){
            tasks[_taskId].supportRequesterWei = 
            tasks[_taskId].supportRequesterWei.add(tasks[_taskId].voterWei[msg.sender]);
        } else if(_vote == 1) {
            tasks[_taskId].supportWorkerWei = 
            tasks[_taskId].supportWorkerWei.add(tasks[_taskId].voterWei[msg.sender]);
        }
    }


    /// @dev Only CEO can fix how many seconds per block is currently observed.
    function setSecondsPerBlock(uint256 secs)
    external
    onlyCEO
    {
        secondsPerBlock = secs;
    }

    /// @dev A payable method for requester to add money for task reward.
    function addMoneyForReward(uint _taskId)
        external
        payable
    {
        require(msg.sender == taskIndexToRequester[_taskId]);
        tasks[_taskId].totalWeiPaid = tasks[_taskId].totalWeiPaid.add(msg.value);
        tasks[_taskId].taskRewardWei = tasks[_taskId].taskRewardWei.add(msg.value);
    }


    function getTaskVoteAddress(uint _taskId)
        public
        view
    returns(address)
    {
        return tasks[_taskId].voteAddr;
    }


    function getVoteFactor(
        uint _taskId
    )
        internal
        view
        returns(address _requester, address _worker)
    {
        _requester = taskIndexToRequester[_taskId];
        _worker = taskIdToWorker[_taskId];
    }
    
    /// @dev the facet of userFee in task, for voter to add votes(Fee).
    function addUserFee(uint _taskId, uint _fee)
        internal
        returns(bool)
    {
        tasks[_taskId].voterFee[msg.sender] = tasks[_taskId].voterFee[msg.sender].add(_fee); 
        return true;
    }
    
    
    /// @dev the facet of userFee in task, for voter to sub votes(Fee)
    function subUserFee(uint _taskId)
        internal
        returns(uint)
    {
        uint _fee = tasks[_taskId].voterFee[msg.sender];
        tasks[_taskId].voterFee[msg.sender] = 0; 
        return _fee;
    }    
    
    /// @dev the facet of userFee in task, for voter to add votes(Wei).
    function addUserWei(uint _taskId, uint _wei)
        internal
        returns(bool)
    {
        tasks[_taskId].voterWei[msg.sender] = tasks[_taskId].voterWei[msg.sender].add(_wei); 
        return true;
    }
    
    
    /// @dev the facet of userFee in task, for voter to sub votes(Wei)
    function subUserWei(uint _taskId)
        internal
        returns(uint)
    {
        uint _wei = tasks[_taskId].voterWei[msg.sender];
        tasks[_taskId].voterWei[msg.sender] = 0;
        return _wei;
    }    
    
    /// @dev voters get their reward
    function getSharedReward(uint _taskId, uint winSide)
        internal
        view
    returns(uint)
    {
        uint _wei = tasks[_taskId].voterWei[msg.sender];
        uint accountForAll;
        uint sharedRewardAll;
        if(winSide == 0){
            // requester win
            accountForAll = tasks[_taskId].supportRequesterWei;
            sharedRewardAll = tasks[_taskId].supportWorkerWei + tasks[_taskId].workerDepositWei;
        }
        else {
            // worker win
            accountForAll = tasks[_taskId].supportWorkerWei;
            sharedRewardAll = tasks[_taskId].supportRequesterWei + tasks[_taskId].requesterDepositWei;
        }
         uint sharedWei = _wei.mul(sharedRewardAll).div(accountForAll);
         return sharedWei;
    }
    
    /// @dev candidate get back reward
    function getTaskReward(uint _taskId , uint _winChoice)
        public
        view
    returns( address , uint)
    {
        if(_winChoice == 0) 
            return (taskIndexToRequester[_taskId], tasks[_taskId].taskRewardWei);
        else if(_winChoice == 1)
            return (taskIdToWorker[_taskId],tasks[_taskId].taskRewardWei);
    }
    
    /*
        struct Task {
        // The task name, that can't change.
        string taskName;
        // Introduction of the task.
        string content;
        // The detail of the task content, that stores hash of the content in IPFS.
        string detailContentHash;
        //Worker's level need to be satisfied, which is the minimum standard for worker.
        uint workerLevel;
        TaskStatus status;      // The task's phase, which is always changing until the task is finish 
        uint totalWeiPaid;      //Total amount paid into the task (in Wei)
        uint taskRewardWei;     //Paid for the task
        string resultHash;      // Worker submit his work, which is stored in IPFS
        uint256 createTime;      // The task Create time.
        uint256 createBlock;     // The task's create block;
        uint requesterDepositWei;  // Requester's deposit for vote
        uint supportRequesterWei; // wei of voter who support requester
        uint workerDepositWei;    // Worker's deposit for vote
        uint supportWorkerWei;
        address voteAddr;       //This is the vote address in ethereum
        string requesterProof;  //requester's proof for voter, it helps voter to make decision
        string workerProof;     //worker's proof for voter, it helps voter to make decision

        mapping(address => uint) voterFee; //User voting cost, Fee transfer from user base
        mapping(address => uint) voterWei; //User voting cost, Wei is Eth unit
    }
    */
    function GetTaskInfo_1(uint _taskId)
        external
        view
    returns(        
        string taskName,
        // Introduction of the task.
        string content,
        // The detail of the task content, that stores hash of the content in IPFS.
        string detailContentHash,
        //Worker's level need to be satisfied, which is the minimum standard for worker.
        uint workerLevel,
        TaskStatus status,      // The task's phase, which is always changing until the task is finish 
        uint totalWeiPaid,      //Total amount paid into the task (in Wei)
        uint taskRewardWei,     //Paid for the task
        string resultHash,      // Worker submit his work, which is stored in IPFS
        uint256 createTime,      // The task Create time.
        uint256 createBlock
        ){
        taskName = tasks[_taskId].taskName;
        content = tasks[_taskId].content;
        detailContentHash = tasks[_taskId].detailContentHash;
        workerLevel = tasks[_taskId].workerLevel;
        status = tasks[_taskId].status;
        totalWeiPaid = tasks[_taskId].totalWeiPaid;
        taskRewardWei = tasks[_taskId].taskRewardWei;
        resultHash = tasks[_taskId].resultHash;
        createTime = tasks[_taskId].createTime;
        createBlock = tasks[_taskId].createBlock;
    }
}

