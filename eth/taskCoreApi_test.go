package eth

import (
	"github.com/helllyf/CrowdSourcing/eth/Watcher"
	"testing"
)

var bc *BcBase

func TestCreateTaskCreate(t *testing.T) {
	bcbase := new(BcBase)
	bc = bcbase
	bcbase.InitEthClient()
	bcbase.DeployTaskContract()
	Watcher.WatcherCreateTask(bcbase.task)
	//注册用户测试
	bcbase.RegisterUser()
	//创建任务测试
	taskName := "任务1"
	content	:= "这是一个测试任务"
	detailContent := "QmSWd9uySbE2KAq92pc1WHwe4Wiv3evtSQzacC63DoJGHC"
	workLevel := 0
	bcbase.CreatTask(taskName,content,detailContent, uint(workLevel))
	bcbase.CreatTask(taskName,content,detailContent, uint(workLevel))
	//<-ch
}

func TestAcceptTask(t *testing.T){
	bc.task.AcceptTask()
}
