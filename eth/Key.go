package eth

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

type EthKey struct{
	account 	string
	privateKey  string
}

var ethKeys = [10]EthKey{
EthKey{"0xe280029a7867ba5c9154434886c241775ea87e53","f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5"},
{"0x68db32d26d9529b2a142927c6f1af248fc6ba7e9","91821f9af458d612362136648fc8552a47d8289c0f25a8a1bf0860510332cef9"},
{"0x35bb6ef95c72bf4804334bb9d6a3c77bef18d81b","bb32062807c162a5243dc9bcf21d8114cb636c376596e1cf2895ec9e5e3e0a68"},
{"0x26d8094a90ad1a4440600fa18073730ef17f5ece","95ce6122165d94aa51b0fcf51021895b39b0ff291aa640c803d5401bd87894d5"},
{"0xb53cc19ad713e00cb71542d064215784c908d387","3af93668029f95d526fc1d2bdefccc120bfe1d26a0462d268e8f6b2f71402ba3"},
{"0xcb98ce2619f90f54052524fb79b03e0261b01bee","3b24a4fdf2e6e1375008c387c5456ce00cb0772435ae1938c2fe833103393b9a"},
{"0x4c4cb54bdbad6c96805b92b8063f3c75b24f65eb","cba858feeb49e1ca8053a5213987a22c3ee83d9f9f396e138940a018dd837ebb"},
{"0x67b27df78bb0f199edbd568e655bd9b2b202866d","df48bfda4cb4b4e094803e923836a8538fbf607da79f6e46d68cdd43fb2f3f88"},
{"0xe93e2b43fc45ccecc056a6ea400972298a304b4b","487efb6249a8a4d45a19c8e5d1e5c7d3f6610a7e69f8f81ddcf368f9a0c0d6d5"},
{"0x288ab710f8dec0b13753fec71161e50ee0cda7e6","bb4cce73db59f456ea427e5862fdb0d5bc038a7d0b930cbb45e1c4f6d122289e"},
}

//ganache-cli -m "much repair shock carbon improve miss forget sock include bullet interest solution"
func GetKey(index int) (string,string,*ecdsa.PrivateKey,common.Address){
	//账户，私钥，ecdsa私钥，公钥
	privateKey := GetPrivateKey(index)
	pubKey := GetPublicKey(index)
	return ethKeys[index].account,ethKeys[index].privateKey,privateKey,pubKey
}

func GetPrivateKey(index int) *ecdsa.PrivateKey{
	privateKey, err := crypto.HexToECDSA(ethKeys[index].privateKey)
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}

func GetPublicKey(index int) common.Address{
	//账户，私钥，ecdsa私钥，公钥
	privateKey := GetPrivateKey(index)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	pubKey := crypto.PubkeyToAddress(*publicKeyECDSA)
	return pubKey
}
