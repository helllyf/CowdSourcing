package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	bc "github.com/ethereum/go-ethereum/ethclient"
	Task "github.com/helllyf/CrowdSourcing/eth/TaskCore"
	"log"
	"math/big"
	"sync"
)

type BcBase struct {
	client 		*bc.Client
	/*交易批次*/
	nonce 		int
	once		sync.Once
	lock  		sync.Mutex

	keyIndex 		int
	/**/
	task	 	*Task.TaskCore
}

func (bcbase *BcBase)InitEthClient(){
	bcbase.once.Do(func(){
		bcbase.keyIndex = 0
		bcbase.nonce = -1
		//fmt.Println(PriK.PublicKey)
		var err error
		//client, err = bc.Dial(ethConf.Url, ethConf.GroupID) //fisco bcos
		bcbase.client, err = bc.Dial("ws://localhost:8545")
		if err != nil {
			log.Fatal(err)
		}
	})
}

func (bcbase *BcBase)getNonce() uint64 {
	bcbase.lock.Lock()
	defer bcbase.lock.Unlock()
	if bcbase.nonce == -1 {
		uintNonce, err := bcbase.client.PendingNonceAt(context.Background(), GetPublicKey(bcbase.keyIndex))
		if err != nil {
			log.Fatal(err)
		}
		bcbase.nonce = int(uintNonce)
	}else{
		bcbase.nonce++
	}

	return uint64(bcbase.nonce)
}

func (bcbase *BcBase)WaitMining(tx *types.Transaction) {
	//fmt.Println("tx is mining")
	ctx := context.Background()
	_, err := bind.WaitMined(ctx, bcbase.client, tx)
	if err != nil {
		log.Fatalf("tx mining error:%v\n", err)
	}
	//fmt.Println("tx is mined")
}
func (bcbase *BcBase)GetAuth() *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(GetPrivateKey(bcbase.keyIndex))
	gasPrice, _ := bcbase.client.SuggestGasPrice(context.Background())
	auth.Nonce 		= big.NewInt(int64(bcbase.getNonce()))
	auth.Value 		= big.NewInt(0) // in wei
	auth.GasLimit 	= uint64(5683526) // in units

	auth.GasPrice 	= gasPrice
	return auth
}

func (bcbase *BcBase)GetBind()  *bind.CallOpts{
	return &bind.CallOpts{From: GetPublicKey(bcbase.keyIndex)}
}

func (bcbase *BcBase)GetBindFilter() *bind.FilterOpts{
	return &bind.FilterOpts{Context:context.Background(),Start:0,End:nil}
}
func (bcbase *BcBase)GetBindWatcher() *bind.WatchOpts{
	return &bind.WatchOpts{Context:context.Background()}
}
func (bcbase *BcBase)ChangeKey(keyIndex int) {
	bcbase.keyIndex = keyIndex
}
