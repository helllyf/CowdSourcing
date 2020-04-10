package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	Task "github.com/helllyf/CrowdSourcing/eth/TaskCore"
	"log"
	"math/big"
)

func (bc *BcBase)DeployTaskContract() common.Address{
	taskAddr, tx, _task, err := Task.DeployTaskCore(bc.GetAuth(),bc.client)
	if err != nil {
		log.Fatal(err)
	}
	bc.WaitMining(tx)
	bc.task = _task

	return taskAddr
}

func (bc *BcBase)ListenEvent(){
	/*
	iter,err := bc.task.FilterCreateTask(bc.GetBindFilter())
	if err != nil{
		log.Fatal(err)
	}
	for iter.Next(){
		fmt.Println(iter.Event.Creator.String())
		fmt.Println(iter.Event.TaskId.String())
	}
*/
	sink := make(chan *Task.TaskCoreCreateTask)
	sub,err := bc.task.WatchCreateTask(bc.GetBindWatcher(),sink)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case createEvent := <-sink:
			fmt.Println(createEvent.Creator.String())
			fmt.Println(createEvent.TaskId.String())
		}
	}

}

func (bc *BcBase)RegisterUser(){
	tx,ok := bc.task.RegisterUser(bc.GetAuth())
	if ok != nil {
		log.Fatal(ok)
	}
	bc.WaitMining(tx)
}

func (bc *BcBase)User() {
	user,err := bc.task.Users(bc.GetBind(),GetPublicKey(bc.keyIndex))
	if err != nil{
		log.Fatal(err)
	}
	_ = user
	//todo user's big int to int
}

func (bc *BcBase) CreatTask(
	_taskName string,
	_content string,
	_detailContentHash string,
	_workerLevel uint,
	) {
	tx,err := bc.task.CreateTask(bc.GetAuth(),_taskName,_content,_detailContentHash,big.NewInt(int64(_workerLevel)))
	fmt.Println("create task")
	if err != nil {
		log.Fatal(err)
	}
	bc.WaitMining(tx)
}

func (bc *BcBase) AcceptTask(taskId uint) {
	tx,err := bc.task.AcceptTask(bc.GetAuth(),big.NewInt(int64(taskId)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("accept task")
	bc.WaitMining(tx)
}





