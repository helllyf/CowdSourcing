package Watcher

import (
	"fmt"
	"github.com/helllyf/CrowdSourcing/eth/TaskCore"
	"log"
)

func WatcherCreateTask(contract *TaskCore.TaskCore)  {
	ch := make(chan *TaskCore.TaskCoreCreateTask)
	sub,err := contract.WatchCreateTask(nil, ch)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case err := <-sub.Err():
				fmt.Errorf(err.Error())
			case log := <-ch:
				fmt.Printf("[Create event]\n")
				fmt.Printf("From: %s\n", log.Creator.String())
				fmt.Printf("TaskId: %d\n", log.TaskId.Int64())
			}
		}
	}()
	fmt.Println("Event watching...")
	//<-ch
}
