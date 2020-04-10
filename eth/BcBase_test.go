package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"testing"
)

func TestInit(t *testing.T)  {
	bcbase := new(BcBase)
	bcbase.InitEthClient()
	account := common.HexToAddress("0xe280029a7867ba5c9154434886c241775ea87e53")
	balance, err := bcbase.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034
}

func TestCreateTaskContract(t *testing.T) {
	bcbase := new(BcBase)
	bcbase.InitEthClient()
	taskAddr := bcbase.DeployTaskContract()
	fmt.Println(taskAddr.Hex())
}


