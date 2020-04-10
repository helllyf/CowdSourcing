pragma solidity >=0.4.21 <0.7.0;

import "./util/SafeMath.sol";
import "./TaskCore.sol";
import  "./User.sol";

/**
@title CrowdSourcing
@dev Allows users to create tasks by paying ETH, and
other users can get paid for completing tasks. If the
task result is not satisfactory, the task creator would
start a poll that USES the wisdom of the masses to
determine if the task is actually done.
*/
contract CrowdSourcing {

    using SafeMath256 for uint;
    using SafeMath32 for uint32;
    using SafeMath8 for uint8;

    // ==========
    // DATA STRUCTURE AND STATE VARIABLES:
    // ==========

    constructor ()
        public
     {
     }
    
    /**
     * description: Release Task By Requester
     * params: taskName
     * return:
     **/
    function releaseTask (
        string memory _taskName,
        string memory _content,
        string memory _detailContentHash,
        uint _workerLevel
    )
    public
    {

    }

    /**
     * description: Get Task address by task id
     * params: _taskId
     * return:
     **/
    function getTask(uint _taskId)
        public
        //returns(address)
    {
        //return taskPool[_taskId];
    }

    /**
     * description: Worker takes part in task
     * params: _taskId
     * return:
     **/
    function acceptTask(uint _taskId)
        public
        returns(bool)
    {
        Task task = taskPool[_taskId];
        bool isAcceptable = task.acceptTask(msg.sender);
        if( isAcceptable == false)
            return false;


        return true;
    }
    
    /**
     * description: This function can only be executed after
       the task has been completed or after the vote has
       been completed.
     * params:
     * return:
     **/
     function finishTask()
        public
     {
     }

}
