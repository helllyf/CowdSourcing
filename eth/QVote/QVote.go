// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package QVote

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// QVoteABI is the input ABI used to generate the binding from.
const QVoteABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"commitment\",\"type\":\"bytes32\"},{\"name\":\"_numVotes\",\"type\":\"uint32\"}],\"name\":\"commitVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sortitionRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"toggleContractActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"randomMod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"firstVoteCost\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRandomNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"},{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"verifySignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rateGrowStep\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_vote\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"revealVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voterCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_numVotes\",\"type\":\"uint256\"},{\"name\":\"_vote\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"hash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"verifyTest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_registerTime\",\"type\":\"uint256\"},{\"name\":\"_honor\",\"type\":\"uint256\"},{\"name\":\"_numVotes\",\"type\":\"uint256\"}],\"name\":\"commitVoteAux\",\"outputs\":[{\"name\":\"_isChoice\",\"type\":\"bool\"},{\"name\":\"_voteCost\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_userId\",\"type\":\"uint256\"}],\"name\":\"uploadHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_factor\",\"type\":\"uint256\"},{\"name\":\"_userId\",\"type\":\"uint256\"}],\"name\":\"uploadFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"intervalBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"completeVote\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"voterToCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterResult\",\"outputs\":[{\"name\":\"isVoter\",\"type\":\"bool\"},{\"name\":\"isWinSide\",\"type\":\"bool\"},{\"name\":\"winChoice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"h\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"recoverSigner2\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voterSelectedNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"name\":\"numVotes\",\"type\":\"uint32\"},{\"name\":\"commitment\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isVoteCompleted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"voteCommitPhase\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_requester\",\"type\":\"address\"},{\"name\":\"_worker\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_closeTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_description\",\"type\":\"bytes32\"}],\"name\":\"VoteCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_description\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_closeTime\",\"type\":\"uint256\"}],\"name\":\"VoterApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"topChoice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"VoteCompleted\",\"type\":\"event\"}]"

// QVoteBin is the compiled bytecode used for deploying new contracts.
var QVoteBin = "0x6080604052606460018190556002556005805460a860020a60ff02191690556000600c553480156200003057600080fd5b5060405160408062001d908339810160405280516020909101516000805433600160a060020a03199182168117909255600580549091169091179055600160a060020a03821615156200008257600080fd5b600160a060020a03811615156200009857600080fd5b620000cd7f72657175657374657200000000000000000000000000000000000000000000008364010000000062000181810204565b50620001037f776f726b657200000000000000000000000000000000000000000000000000008264010000000062000181810204565b50506007805443600655600a600f556005601081905563ffffffff199091164263ffffffff1617604060020a63ffffffff021916680a00000000000000001790915560118190556009805461ff0019169055805460a060020a60ff0219167401000000000000000000000000000000000000000017905550620004fe565b60055460009081907501000000000000000000000000000000000000000000900460ff1615620001b057600080fd5b60095460ff908116106200024b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f43616e206e6f7420616464206d6f7265207468616e203235352063616e64696460448201527f617465732e000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5060095460ff166200026d8160016401000000006200183e620002cc82021704565b60ff81166000818152600a60205260409020600181018790556002018054600160a060020a031916600160a060020a0387161790556009805460ff191690911790559050620002c583640100000000620002ec810204565b5092915050565b600082820160ff8085169082161015620002e557600080fd5b9392505050565b620002f66200042c565b506040805160a081018252600160a060020a038381168252600060208084018281528486018381528651808401909752838752606086019687526080860184905260048054600181018083559190955286517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b60059096029586018054600160a060020a0319169190971617865591517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c850155517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19d840155945180519495948694936200040b937f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19e90910192019062000459565b50608091909101516004909101805460ff1916911515919091179055505050565b6040805160a081018252600080825260208201819052918101829052606080820152608081019190915290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200049c57805160ff1916838001178555620004cc565b82800160010185558215620004cc579182015b82811115620004cc578251825591602001919060010190620004af565b50620004da929150620004de565b5090565b620004fb91905b80821115620004da5760008155600101620004e5565b90565b611882806200050e6000396000f3006080604052600436106101745763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166303b8a89f811461017957806308731651146101a85780631385d24c146101cf57806316c16c20146101e45780631ea77adb146101f957806325b814f41461022757806331cd41991461023c57806335aa946c146102b95780633932d6b7146102ce57806342169e48146102f55780635980e67f1461030a5780636e23c5151461032857806383197ef014610349578063837833581461035e5780638bb89e49146103a35780638da5cb5b146103ca5780638db372b7146103fb5780638f54be0e1461041657806397aba7f91461042b5780639b17a9a114610489578063a567dd751461049e578063af2287e3146104bf578063bdb33d4814610500578063c7fb60e514610524578063ccebafd814610539578063d8bff5a51461054e578063e0f721271461058f578063f2fde38b146105a4578063fcd31320146105c5575b600080fd5b34801561018557600080fd5b506101a6600160a060020a036004351660243563ffffffff604435166105da565b005b3480156101b457600080fd5b506101bd6107e1565b60408051918252519081900360200190f35b3480156101db57600080fd5b506101a66107e7565b3480156101f057600080fd5b506101bd61082c565b34801561020557600080fd5b5061020e610832565b6040805163ffffffff9092168252519081900360200190f35b34801561023357600080fd5b506101bd61084a565b34801561024857600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526102a595833595369560449491939091019190819084018382808284375094975050509235600160a060020a0316935061099f92505050565b604080519115158252519081900360200190f35b3480156102c557600080fd5b506101bd6109c5565b3480156102da57600080fd5b506101a6600160a060020a03600435166024356044356109cb565b34801561030157600080fd5b506101bd610d4f565b34801561031657600080fd5b506101bd600435602435604435610d55565b34801561033457600080fd5b506102a5600160a060020a0360043516610e69565b34801561035557600080fd5b506102a5610f03565b34801561036a57600080fd5b50610388600160a060020a0360043516602435604435606435610f21565b60408051921515835260208301919091528051918290030190f35b3480156103af57600080fd5b506101a6600480359060248035908101910135604435611086565b3480156103d657600080fd5b506103df61114b565b60408051600160a060020a039092168252519081900360200190f35b34801561040757600080fd5b506102a560043560243561115a565b34801561042257600080fd5b506101bd611284565b34801561043757600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526103df95833595369560449491939091019190819084018382808284375094975061128a9650505050505050565b34801561049557600080fd5b506101bd611391565b3480156104aa57600080fd5b506101bd600160a060020a036004351661144f565b3480156104cb57600080fd5b506104e0600160a060020a0360043516611461565b604080519315158452911515602084015282820152519081900360600190f35b34801561050c57600080fd5b506103df60043560ff602435166044356064356114bf565b34801561053057600080fd5b506101bd611622565b34801561054557600080fd5b506101bd611628565b34801561055a57600080fd5b5061056f600160a060020a036004351661162e565b6040805163ffffffff909316835260208301919091528051918290030190f35b34801561059b57600080fd5b506102a5611650565b3480156105b057600080fd5b506101a6600160a060020a036004351661167e565b3480156105d157600080fd5b506101a66116c4565b6105e261178f565b600554600160a060020a031633146105f957600080fd5b6001600954610100900460ff16600381111561061157fe5b148061063257506002600954610100900460ff16600381111561063057fe5b145b151561063d57600080fd5b6005547501000000000000000000000000000000000000000000900460ff161561066657600080fd5b600160a060020a0384166000908152600b60205260409020805460ff1916600190811790915563ffffffff831610156106e9576040805160e560020a62461bcd02815260206004820152601e60248201527f4d757374207375626d6974206174206c65617374206f6e6520766f74652e0000604482015290519081900360640190fd5b600160a060020a0384166000908152600d602052604090205463ffffffff161561075d576040805160e560020a62461bcd02815260206004820152601b60248201527f43616e206f6e6c7920636f6d6d697420766f746573206f6e63652e0000000000604482015290519081900360640190fd5b5060408051808201825263ffffffff83811682526020808301868152600160a060020a0388166000908152600d9092529390208251815463ffffffff19169083161781559251600193840155600c5491926107b992916116ec16565b600c819055601154116107db576000600c556009805461ff0019166102001790555b50505050565b600f5481565b6005805475ff00000000000000000000000000000000000000000019811675010000000000000000000000000000000000000000009182900460ff1615909102179055565b60015481565b60075468010000000000000000900463ffffffff1681565b600080806001815b60045481101561089457600480548290811061086a57fe5b600091825260209091206004600590920201015460ff16151561088c57600091505b600101610852565b81156109225760009350600092505b600454831015610904576108e06004848154811015156108bf57fe5b906000526020600020906005020160010154856116ec90919063ffffffff16565b93506108f76001548561170590919063ffffffff16565b93506001909201916108a3565b60015461091890859063ffffffff6116ec16565b9350839450610998565b6002546003540143101561093557600080fd5b60009350600092505b60045483101561097e5761095a6004848154811015156108bf57fe5b93506109716001548561170590919063ffffffff16565b935060019092019161093e565b60015461099290859063ffffffff6116ec16565b93508394505b5050505090565b6000806109ac858561128a565b600160a060020a03908116931692909214949350505050565b60105481565b6109d361178f565b60055460009081908190600160a060020a031633146109f157600080fd5b6005547501000000000000000000000000000000000000000000900460ff1615610a1a57600080fd5b336000908152600b602052604090205460ff161515600114610aac576040805160e560020a62461bcd02815260206004820152602360248201527f6d73672e73656e646572206d75737420626520617070726f76656420746f207660448201527f6f74650000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0387166000908152600d602090815260409182902082518084018452815463ffffffff168082526001909201548184015283518084018390528085018b905260608082018b90528551808303909101815260809091019485905280519198509196509092918291908401908083835b60208310610b415780518252601f199092019160209182019101610b22565b51815160209384036101000a6000190180199092169116179052604051919093018190039020918801519195505084149150610bef9050576040805160e560020a62461bcd02815260206004820152602760248201527f4861736820646f6573206e6f74206d617463682074686520636f6d6d6974746560448201527f6420686173682e00000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0387166000908152600d60205260408120805463ffffffff1916815560010181905586118015610c2b575060095460ff168611155b1515610ca7576040805160e560020a62461bcd02815260206004820152602a60248201527f566f7465206d75737420626520706c6163656420666f7220612076616c69642060448201527f63616e6469646174652e00000000000000000000000000000000000000000000606482015290519081900360840190fd5b506000858152600a602052604090205463ffffffff90811690610cd0908290858116906116ec16565b6000878152600a60209081526040808320805463ffffffff191663ffffffff86811691909117909155600160a060020a038c168452600e909252909120889055600c54919250610d2491906001906116ec16565b600c81905560115411610d46576000600c556009805461ff0019166103001790555b50505050505050565b600c5481565b60095460009060ff16831115610ddb576040805160e560020a62461bcd02815260206004820152602a60248201527f566f7465206d75737420626520706c6163656420666f7220612076616c69642060448201527f63616e6469646174652e00000000000000000000000000000000000000000000606482015290519081900360840190fd5b604080516020808201879052818301869052606080830186905283518084039091018152608090920192839052815191929182918401908083835b60208310610e355780518252601f199092019160209182019101610e16565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120979650505050505050565b60408051608081018252604181527f39621d9bcfd838bf3f4666af7b33c355d61d9102f394aca137acdf7bd0c20e5260208201527f1af2df88a03f30121993e7c591c7948e3da559e64b1d2d8355094479e3a72a5391810191909152600060608201819052907f4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45610efb81838661099f565b949350505050565b60008054600160a060020a03163314610f1b57600080fd5b33ff5b90565b6005546000908190819081908190819081908190600160a060020a03163314610f4957600080fd5b60075463ffffffff168b1115610f655760009750879650611077565b610f6d61084a565b95508b866040516020018083600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b60208310610fe75780518252601f199092019160209182019101610fc8565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912097508892508791505081151561102257fe5b0693506006544303925060646010548481151561103b57fe5b04600f548c0101870281151561104d57fe5b0491508184101561106f5761106189611726565b905060018197509750611077565b600097508796505b50505050505094509492505050565b6110eb8484848080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050506004848154811015156110cc57fe5b6000918252602090912060059091020154600160a060020a031661099f565b156107db578360048281548110151561110057fe5b600091825260209091206002600590920201015560048054849184918490811061112657fe5b906000526020600020906005020160030191906111449291906117a6565b5050505050565b600554600160a060020a031681565b60008083604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b602083106111ac5780518252601f19909201916020918201910161118d565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090506004838154811015156111ea57fe5b60009182526020909120600260059092020101548114156112785760015461121990859063ffffffff61170516565b600480548590811061122757fe5b906000526020600020906005020160010181905550600160048481548110151561124d57fe5b60009182526020909120600590910201600401805460ff19169115159190911790556001915061127d565b600091505b5092915050565b60025481565b600080600080845160411415156112eb576040805160e560020a62461bcd02815260206004820152601660248201527f5265717569726520636f7272656374206c656e67746800000000000000000000604482015290519081900360640190fd5b50505060208201516040830151606084015160001a601b60ff8216101561131057601b015b8060ff16601b148061132557508060ff16601c145b151561137b576040805160e560020a62461bcd02815260206004820152601b60248201527f5369676e61747572652076657273696f6e206e6f74206d617463680000000000604482015290519081900360640190fd5b611387868285856114bf565b9695505050505050565b600080806003600954610100900460ff1660038111156113ad57fe5b14611402576040805160e560020a62461bcd02815260206004820181905260248201527f766f7465206d75737420626520696e20436f6d706c657465642050686173652e604482015290519081900360640190fd5b506000905060015b60095460ff16811015611449576000828152600a60205260408082205483835291205463ffffffff91821691161115611441578091505b60010161140a565b50919050565b600e6020526000908152604090205481565b600160a060020a0381166000908152600b602052604081205460ff169080611487611391565b600160a060020a0385166000908152600e60205260409020549091508114156114b357600191506114b8565b600091505b9193909250565b604080518082018252601c8082527f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080840191825293516000948593849386938c9301918291908083835b6020831061152b5780518252601f19909201916020918201910161150c565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190819052835193945092839250908401908083835b6020831061158b5780518252601f19909201916020918201910161156c565b51815160209384036101000a600019018019909216911617905260408051929094018290038220600080845283830180875282905260ff8f1684870152606084018e9052608084018d905294519098506001965060a080840196509194601f19820194509281900390910191865af115801561160b573d6000803e3d6000fd5b5050604051601f1901519998505050505050505050565b60115481565b60035481565b600d602052600090815260409020805460019091015463ffffffff9091169082565b60006003600954610100900460ff16600381111561166a57fe5b141561167857506001610f1e565b50600090565b600554600160a060020a0316331461169557600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600554600160a060020a031633146116db57600080fd5b6009805461ff001916610100179055565b6000828201838110156116fe57600080fd5b9392505050565b600081151561171357600080fd5b818381151561171e57fe5b069392505050565b60075460009081906116fe9063ffffffff68010000000000000000909104811690611755908690819061176116565b9063ffffffff61176116565b600080831515611774576000915061127d565b5082820282848281151561178457fe5b04146116fe57600080fd5b604080518082019091526000808252602082015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106117e75782800160ff19823516178555611814565b82800160010185558215611814579182015b828111156118145782358255916020019190600101906117f9565b50611820929150611824565b5090565b610f1e91905b80821115611820576000815560010161182a565b600082820160ff80851690821610156116fe57600080fd00a165627a7a72305820633044effba8e5014091ec7965771142d55f0d4a1afa3b61b24a4cbbff5526ad0029"

// DeployQVote deploys a new Ethereum contract, binding an instance of QVote to it.
func DeployQVote(auth *bind.TransactOpts, backend bind.ContractBackend, _requester common.Address, _worker common.Address) (common.Address, *types.Transaction, *QVote, error) {
	parsed, err := abi.JSON(strings.NewReader(QVoteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(QVoteBin), backend, _requester, _worker)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QVote{QVoteCaller: QVoteCaller{contract: contract}, QVoteTransactor: QVoteTransactor{contract: contract}, QVoteFilterer: QVoteFilterer{contract: contract}}, nil
}

// QVote is an auto generated Go binding around an Ethereum contract.
type QVote struct {
	QVoteCaller     // Read-only binding to the contract
	QVoteTransactor // Write-only binding to the contract
	QVoteFilterer   // Log filterer for contract events
}

// QVoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type QVoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QVoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QVoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QVoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QVoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QVoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QVoteSession struct {
	Contract     *QVote            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QVoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QVoteCallerSession struct {
	Contract *QVoteCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QVoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QVoteTransactorSession struct {
	Contract     *QVoteTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QVoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type QVoteRaw struct {
	Contract *QVote // Generic contract binding to access the raw methods on
}

// QVoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QVoteCallerRaw struct {
	Contract *QVoteCaller // Generic read-only contract binding to access the raw methods on
}

// QVoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QVoteTransactorRaw struct {
	Contract *QVoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQVote creates a new instance of QVote, bound to a specific deployed contract.
func NewQVote(address common.Address, backend bind.ContractBackend) (*QVote, error) {
	contract, err := bindQVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QVote{QVoteCaller: QVoteCaller{contract: contract}, QVoteTransactor: QVoteTransactor{contract: contract}, QVoteFilterer: QVoteFilterer{contract: contract}}, nil
}

// NewQVoteCaller creates a new read-only instance of QVote, bound to a specific deployed contract.
func NewQVoteCaller(address common.Address, caller bind.ContractCaller) (*QVoteCaller, error) {
	contract, err := bindQVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QVoteCaller{contract: contract}, nil
}

// NewQVoteTransactor creates a new write-only instance of QVote, bound to a specific deployed contract.
func NewQVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*QVoteTransactor, error) {
	contract, err := bindQVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QVoteTransactor{contract: contract}, nil
}

// NewQVoteFilterer creates a new log filterer instance of QVote, bound to a specific deployed contract.
func NewQVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*QVoteFilterer, error) {
	contract, err := bindQVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QVoteFilterer{contract: contract}, nil
}

// bindQVote binds a generic wrapper to an already deployed contract.
func bindQVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QVoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QVote *QVoteRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _QVote.Contract.QVoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QVote *QVoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QVote.Contract.QVoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QVote *QVoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QVote.Contract.QVoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QVote *QVoteCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _QVote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QVote *QVoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QVote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QVote *QVoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QVote.Contract.contract.Transact(opts, method, params...)
}

// CommitVoteAux is a free data retrieval call binding the contract method 0x83783358.
//
// Solidity: function commitVoteAux(address _account, uint256 _registerTime, uint256 _honor, uint256 _numVotes) constant returns(bool _isChoice, uint256 _voteCost)
func (_QVote *QVoteCaller) CommitVoteAux(opts *bind.CallOpts, _account common.Address, _registerTime *big.Int, _honor *big.Int, _numVotes *big.Int) (struct {
	IsChoice bool
	VoteCost *big.Int
}, error) {
	ret := new(struct {
		IsChoice bool
		VoteCost *big.Int
	})
	out := ret
	err := _QVote.contract.Call(opts, out, "commitVoteAux", _account, _registerTime, _honor, _numVotes)
	return *ret, err
}

// CommitVoteAux is a free data retrieval call binding the contract method 0x83783358.
//
// Solidity: function commitVoteAux(address _account, uint256 _registerTime, uint256 _honor, uint256 _numVotes) constant returns(bool _isChoice, uint256 _voteCost)
func (_QVote *QVoteSession) CommitVoteAux(_account common.Address, _registerTime *big.Int, _honor *big.Int, _numVotes *big.Int) (struct {
	IsChoice bool
	VoteCost *big.Int
}, error) {
	return _QVote.Contract.CommitVoteAux(&_QVote.CallOpts, _account, _registerTime, _honor, _numVotes)
}

// CommitVoteAux is a free data retrieval call binding the contract method 0x83783358.
//
// Solidity: function commitVoteAux(address _account, uint256 _registerTime, uint256 _honor, uint256 _numVotes) constant returns(bool _isChoice, uint256 _voteCost)
func (_QVote *QVoteCallerSession) CommitVoteAux(_account common.Address, _registerTime *big.Int, _honor *big.Int, _numVotes *big.Int) (struct {
	IsChoice bool
	VoteCost *big.Int
}, error) {
	return _QVote.Contract.CommitVoteAux(&_QVote.CallOpts, _account, _registerTime, _honor, _numVotes)
}

// CompleteVote is a free data retrieval call binding the contract method 0x9b17a9a1.
//
// Solidity: function completeVote() constant returns(uint256)
func (_QVote *QVoteCaller) CompleteVote(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "completeVote")
	return *ret0, err
}

// CompleteVote is a free data retrieval call binding the contract method 0x9b17a9a1.
//
// Solidity: function completeVote() constant returns(uint256)
func (_QVote *QVoteSession) CompleteVote() (*big.Int, error) {
	return _QVote.Contract.CompleteVote(&_QVote.CallOpts)
}

// CompleteVote is a free data retrieval call binding the contract method 0x9b17a9a1.
//
// Solidity: function completeVote() constant returns(uint256)
func (_QVote *QVoteCallerSession) CompleteVote() (*big.Int, error) {
	return _QVote.Contract.CompleteVote(&_QVote.CallOpts)
}

// CreateBlock is a free data retrieval call binding the contract method 0xccebafd8.
//
// Solidity: function createBlock() constant returns(uint256)
func (_QVote *QVoteCaller) CreateBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "createBlock")
	return *ret0, err
}

// CreateBlock is a free data retrieval call binding the contract method 0xccebafd8.
//
// Solidity: function createBlock() constant returns(uint256)
func (_QVote *QVoteSession) CreateBlock() (*big.Int, error) {
	return _QVote.Contract.CreateBlock(&_QVote.CallOpts)
}

// CreateBlock is a free data retrieval call binding the contract method 0xccebafd8.
//
// Solidity: function createBlock() constant returns(uint256)
func (_QVote *QVoteCallerSession) CreateBlock() (*big.Int, error) {
	return _QVote.Contract.CreateBlock(&_QVote.CallOpts)
}

// FirstVoteCost is a free data retrieval call binding the contract method 0x1ea77adb.
//
// Solidity: function firstVoteCost() constant returns(uint32)
func (_QVote *QVoteCaller) FirstVoteCost(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "firstVoteCost")
	return *ret0, err
}

// FirstVoteCost is a free data retrieval call binding the contract method 0x1ea77adb.
//
// Solidity: function firstVoteCost() constant returns(uint32)
func (_QVote *QVoteSession) FirstVoteCost() (uint32, error) {
	return _QVote.Contract.FirstVoteCost(&_QVote.CallOpts)
}

// FirstVoteCost is a free data retrieval call binding the contract method 0x1ea77adb.
//
// Solidity: function firstVoteCost() constant returns(uint32)
func (_QVote *QVoteCallerSession) FirstVoteCost() (uint32, error) {
	return _QVote.Contract.FirstVoteCost(&_QVote.CallOpts)
}

// GetRandomNum is a free data retrieval call binding the contract method 0x25b814f4.
//
// Solidity: function getRandomNum() constant returns(uint256)
func (_QVote *QVoteCaller) GetRandomNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "getRandomNum")
	return *ret0, err
}

// GetRandomNum is a free data retrieval call binding the contract method 0x25b814f4.
//
// Solidity: function getRandomNum() constant returns(uint256)
func (_QVote *QVoteSession) GetRandomNum() (*big.Int, error) {
	return _QVote.Contract.GetRandomNum(&_QVote.CallOpts)
}

// GetRandomNum is a free data retrieval call binding the contract method 0x25b814f4.
//
// Solidity: function getRandomNum() constant returns(uint256)
func (_QVote *QVoteCallerSession) GetRandomNum() (*big.Int, error) {
	return _QVote.Contract.GetRandomNum(&_QVote.CallOpts)
}

// GetVoterResult is a free data retrieval call binding the contract method 0xaf2287e3.
//
// Solidity: function getVoterResult(address _voter) constant returns(bool isVoter, bool isWinSide, uint256 winChoice)
func (_QVote *QVoteCaller) GetVoterResult(opts *bind.CallOpts, _voter common.Address) (struct {
	IsVoter   bool
	IsWinSide bool
	WinChoice *big.Int
}, error) {
	ret := new(struct {
		IsVoter   bool
		IsWinSide bool
		WinChoice *big.Int
	})
	out := ret
	err := _QVote.contract.Call(opts, out, "getVoterResult", _voter)
	return *ret, err
}

// GetVoterResult is a free data retrieval call binding the contract method 0xaf2287e3.
//
// Solidity: function getVoterResult(address _voter) constant returns(bool isVoter, bool isWinSide, uint256 winChoice)
func (_QVote *QVoteSession) GetVoterResult(_voter common.Address) (struct {
	IsVoter   bool
	IsWinSide bool
	WinChoice *big.Int
}, error) {
	return _QVote.Contract.GetVoterResult(&_QVote.CallOpts, _voter)
}

// GetVoterResult is a free data retrieval call binding the contract method 0xaf2287e3.
//
// Solidity: function getVoterResult(address _voter) constant returns(bool isVoter, bool isWinSide, uint256 winChoice)
func (_QVote *QVoteCallerSession) GetVoterResult(_voter common.Address) (struct {
	IsVoter   bool
	IsWinSide bool
	WinChoice *big.Int
}, error) {
	return _QVote.Contract.GetVoterResult(&_QVote.CallOpts, _voter)
}

// Hash is a free data retrieval call binding the contract method 0x5980e67f.
//
// Solidity: function hash(uint256 _numVotes, uint256 _vote, bytes32 salt) constant returns(bytes32)
func (_QVote *QVoteCaller) Hash(opts *bind.CallOpts, _numVotes *big.Int, _vote *big.Int, salt [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "hash", _numVotes, _vote, salt)
	return *ret0, err
}

// Hash is a free data retrieval call binding the contract method 0x5980e67f.
//
// Solidity: function hash(uint256 _numVotes, uint256 _vote, bytes32 salt) constant returns(bytes32)
func (_QVote *QVoteSession) Hash(_numVotes *big.Int, _vote *big.Int, salt [32]byte) ([32]byte, error) {
	return _QVote.Contract.Hash(&_QVote.CallOpts, _numVotes, _vote, salt)
}

// Hash is a free data retrieval call binding the contract method 0x5980e67f.
//
// Solidity: function hash(uint256 _numVotes, uint256 _vote, bytes32 salt) constant returns(bytes32)
func (_QVote *QVoteCallerSession) Hash(_numVotes *big.Int, _vote *big.Int, salt [32]byte) ([32]byte, error) {
	return _QVote.Contract.Hash(&_QVote.CallOpts, _numVotes, _vote, salt)
}

// IntervalBlock is a free data retrieval call binding the contract method 0x8f54be0e.
//
// Solidity: function intervalBlock() constant returns(uint256)
func (_QVote *QVoteCaller) IntervalBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "intervalBlock")
	return *ret0, err
}

// IntervalBlock is a free data retrieval call binding the contract method 0x8f54be0e.
//
// Solidity: function intervalBlock() constant returns(uint256)
func (_QVote *QVoteSession) IntervalBlock() (*big.Int, error) {
	return _QVote.Contract.IntervalBlock(&_QVote.CallOpts)
}

// IntervalBlock is a free data retrieval call binding the contract method 0x8f54be0e.
//
// Solidity: function intervalBlock() constant returns(uint256)
func (_QVote *QVoteCallerSession) IntervalBlock() (*big.Int, error) {
	return _QVote.Contract.IntervalBlock(&_QVote.CallOpts)
}

// IsVoteCompleted is a free data retrieval call binding the contract method 0xe0f72127.
//
// Solidity: function isVoteCompleted() constant returns(bool)
func (_QVote *QVoteCaller) IsVoteCompleted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "isVoteCompleted")
	return *ret0, err
}

// IsVoteCompleted is a free data retrieval call binding the contract method 0xe0f72127.
//
// Solidity: function isVoteCompleted() constant returns(bool)
func (_QVote *QVoteSession) IsVoteCompleted() (bool, error) {
	return _QVote.Contract.IsVoteCompleted(&_QVote.CallOpts)
}

// IsVoteCompleted is a free data retrieval call binding the contract method 0xe0f72127.
//
// Solidity: function isVoteCompleted() constant returns(bool)
func (_QVote *QVoteCallerSession) IsVoteCompleted() (bool, error) {
	return _QVote.Contract.IsVoteCompleted(&_QVote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_QVote *QVoteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_QVote *QVoteSession) Owner() (common.Address, error) {
	return _QVote.Contract.Owner(&_QVote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_QVote *QVoteCallerSession) Owner() (common.Address, error) {
	return _QVote.Contract.Owner(&_QVote.CallOpts)
}

// RandomMod is a free data retrieval call binding the contract method 0x16c16c20.
//
// Solidity: function randomMod() constant returns(uint256)
func (_QVote *QVoteCaller) RandomMod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "randomMod")
	return *ret0, err
}

// RandomMod is a free data retrieval call binding the contract method 0x16c16c20.
//
// Solidity: function randomMod() constant returns(uint256)
func (_QVote *QVoteSession) RandomMod() (*big.Int, error) {
	return _QVote.Contract.RandomMod(&_QVote.CallOpts)
}

// RandomMod is a free data retrieval call binding the contract method 0x16c16c20.
//
// Solidity: function randomMod() constant returns(uint256)
func (_QVote *QVoteCallerSession) RandomMod() (*big.Int, error) {
	return _QVote.Contract.RandomMod(&_QVote.CallOpts)
}

// RateGrowStep is a free data retrieval call binding the contract method 0x35aa946c.
//
// Solidity: function rateGrowStep() constant returns(uint256)
func (_QVote *QVoteCaller) RateGrowStep(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "rateGrowStep")
	return *ret0, err
}

// RateGrowStep is a free data retrieval call binding the contract method 0x35aa946c.
//
// Solidity: function rateGrowStep() constant returns(uint256)
func (_QVote *QVoteSession) RateGrowStep() (*big.Int, error) {
	return _QVote.Contract.RateGrowStep(&_QVote.CallOpts)
}

// RateGrowStep is a free data retrieval call binding the contract method 0x35aa946c.
//
// Solidity: function rateGrowStep() constant returns(uint256)
func (_QVote *QVoteCallerSession) RateGrowStep() (*big.Int, error) {
	return _QVote.Contract.RateGrowStep(&_QVote.CallOpts)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 hash, bytes sig) constant returns(address)
func (_QVote *QVoteCaller) RecoverSigner(opts *bind.CallOpts, hash [32]byte, sig []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "recoverSigner", hash, sig)
	return *ret0, err
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 hash, bytes sig) constant returns(address)
func (_QVote *QVoteSession) RecoverSigner(hash [32]byte, sig []byte) (common.Address, error) {
	return _QVote.Contract.RecoverSigner(&_QVote.CallOpts, hash, sig)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 hash, bytes sig) constant returns(address)
func (_QVote *QVoteCallerSession) RecoverSigner(hash [32]byte, sig []byte) (common.Address, error) {
	return _QVote.Contract.RecoverSigner(&_QVote.CallOpts, hash, sig)
}

// RecoverSigner2 is a free data retrieval call binding the contract method 0xbdb33d48.
//
// Solidity: function recoverSigner2(bytes32 h, uint8 v, bytes32 r, bytes32 s) constant returns(address)
func (_QVote *QVoteCaller) RecoverSigner2(opts *bind.CallOpts, h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "recoverSigner2", h, v, r, s)
	return *ret0, err
}

// RecoverSigner2 is a free data retrieval call binding the contract method 0xbdb33d48.
//
// Solidity: function recoverSigner2(bytes32 h, uint8 v, bytes32 r, bytes32 s) constant returns(address)
func (_QVote *QVoteSession) RecoverSigner2(h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _QVote.Contract.RecoverSigner2(&_QVote.CallOpts, h, v, r, s)
}

// RecoverSigner2 is a free data retrieval call binding the contract method 0xbdb33d48.
//
// Solidity: function recoverSigner2(bytes32 h, uint8 v, bytes32 r, bytes32 s) constant returns(address)
func (_QVote *QVoteCallerSession) RecoverSigner2(h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _QVote.Contract.RecoverSigner2(&_QVote.CallOpts, h, v, r, s)
}

// SortitionRate is a free data retrieval call binding the contract method 0x08731651.
//
// Solidity: function sortitionRate() constant returns(uint256)
func (_QVote *QVoteCaller) SortitionRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "sortitionRate")
	return *ret0, err
}

// SortitionRate is a free data retrieval call binding the contract method 0x08731651.
//
// Solidity: function sortitionRate() constant returns(uint256)
func (_QVote *QVoteSession) SortitionRate() (*big.Int, error) {
	return _QVote.Contract.SortitionRate(&_QVote.CallOpts)
}

// SortitionRate is a free data retrieval call binding the contract method 0x08731651.
//
// Solidity: function sortitionRate() constant returns(uint256)
func (_QVote *QVoteCallerSession) SortitionRate() (*big.Int, error) {
	return _QVote.Contract.SortitionRate(&_QVote.CallOpts)
}

// VerifySignature is a free data retrieval call binding the contract method 0x31cd4199.
//
// Solidity: function verifySignature(bytes32 hash, bytes signature, address signer) constant returns(bool)
func (_QVote *QVoteCaller) VerifySignature(opts *bind.CallOpts, hash [32]byte, signature []byte, signer common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "verifySignature", hash, signature, signer)
	return *ret0, err
}

// VerifySignature is a free data retrieval call binding the contract method 0x31cd4199.
//
// Solidity: function verifySignature(bytes32 hash, bytes signature, address signer) constant returns(bool)
func (_QVote *QVoteSession) VerifySignature(hash [32]byte, signature []byte, signer common.Address) (bool, error) {
	return _QVote.Contract.VerifySignature(&_QVote.CallOpts, hash, signature, signer)
}

// VerifySignature is a free data retrieval call binding the contract method 0x31cd4199.
//
// Solidity: function verifySignature(bytes32 hash, bytes signature, address signer) constant returns(bool)
func (_QVote *QVoteCallerSession) VerifySignature(hash [32]byte, signature []byte, signer common.Address) (bool, error) {
	return _QVote.Contract.VerifySignature(&_QVote.CallOpts, hash, signature, signer)
}

// VerifyTest is a free data retrieval call binding the contract method 0x6e23c515.
//
// Solidity: function verifyTest(address signer) constant returns(bool)
func (_QVote *QVoteCaller) VerifyTest(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "verifyTest", signer)
	return *ret0, err
}

// VerifyTest is a free data retrieval call binding the contract method 0x6e23c515.
//
// Solidity: function verifyTest(address signer) constant returns(bool)
func (_QVote *QVoteSession) VerifyTest(signer common.Address) (bool, error) {
	return _QVote.Contract.VerifyTest(&_QVote.CallOpts, signer)
}

// VerifyTest is a free data retrieval call binding the contract method 0x6e23c515.
//
// Solidity: function verifyTest(address signer) constant returns(bool)
func (_QVote *QVoteCallerSession) VerifyTest(signer common.Address) (bool, error) {
	return _QVote.Contract.VerifyTest(&_QVote.CallOpts, signer)
}

// VoterCount is a free data retrieval call binding the contract method 0x42169e48.
//
// Solidity: function voterCount() constant returns(uint256)
func (_QVote *QVoteCaller) VoterCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "voterCount")
	return *ret0, err
}

// VoterCount is a free data retrieval call binding the contract method 0x42169e48.
//
// Solidity: function voterCount() constant returns(uint256)
func (_QVote *QVoteSession) VoterCount() (*big.Int, error) {
	return _QVote.Contract.VoterCount(&_QVote.CallOpts)
}

// VoterCount is a free data retrieval call binding the contract method 0x42169e48.
//
// Solidity: function voterCount() constant returns(uint256)
func (_QVote *QVoteCallerSession) VoterCount() (*big.Int, error) {
	return _QVote.Contract.VoterCount(&_QVote.CallOpts)
}

// VoterSelectedNum is a free data retrieval call binding the contract method 0xc7fb60e5.
//
// Solidity: function voterSelectedNum() constant returns(uint256)
func (_QVote *QVoteCaller) VoterSelectedNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "voterSelectedNum")
	return *ret0, err
}

// VoterSelectedNum is a free data retrieval call binding the contract method 0xc7fb60e5.
//
// Solidity: function voterSelectedNum() constant returns(uint256)
func (_QVote *QVoteSession) VoterSelectedNum() (*big.Int, error) {
	return _QVote.Contract.VoterSelectedNum(&_QVote.CallOpts)
}

// VoterSelectedNum is a free data retrieval call binding the contract method 0xc7fb60e5.
//
// Solidity: function voterSelectedNum() constant returns(uint256)
func (_QVote *QVoteCallerSession) VoterSelectedNum() (*big.Int, error) {
	return _QVote.Contract.VoterSelectedNum(&_QVote.CallOpts)
}

// VoterToCandidate is a free data retrieval call binding the contract method 0xa567dd75.
//
// Solidity: function voterToCandidate(address ) constant returns(uint256)
func (_QVote *QVoteCaller) VoterToCandidate(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _QVote.contract.Call(opts, out, "voterToCandidate", arg0)
	return *ret0, err
}

// VoterToCandidate is a free data retrieval call binding the contract method 0xa567dd75.
//
// Solidity: function voterToCandidate(address ) constant returns(uint256)
func (_QVote *QVoteSession) VoterToCandidate(arg0 common.Address) (*big.Int, error) {
	return _QVote.Contract.VoterToCandidate(&_QVote.CallOpts, arg0)
}

// VoterToCandidate is a free data retrieval call binding the contract method 0xa567dd75.
//
// Solidity: function voterToCandidate(address ) constant returns(uint256)
func (_QVote *QVoteCallerSession) VoterToCandidate(arg0 common.Address) (*big.Int, error) {
	return _QVote.Contract.VoterToCandidate(&_QVote.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(uint32 numVotes, bytes32 commitment)
func (_QVote *QVoteCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (struct {
	NumVotes   uint32
	Commitment [32]byte
}, error) {
	ret := new(struct {
		NumVotes   uint32
		Commitment [32]byte
	})
	out := ret
	err := _QVote.contract.Call(opts, out, "votes", arg0)
	return *ret, err
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(uint32 numVotes, bytes32 commitment)
func (_QVote *QVoteSession) Votes(arg0 common.Address) (struct {
	NumVotes   uint32
	Commitment [32]byte
}, error) {
	return _QVote.Contract.Votes(&_QVote.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(uint32 numVotes, bytes32 commitment)
func (_QVote *QVoteCallerSession) Votes(arg0 common.Address) (struct {
	NumVotes   uint32
	Commitment [32]byte
}, error) {
	return _QVote.Contract.Votes(&_QVote.CallOpts, arg0)
}

// CommitVote is a paid mutator transaction binding the contract method 0x03b8a89f.
//
// Solidity: function commitVote(address _voter, bytes32 commitment, uint32 _numVotes) returns()
func (_QVote *QVoteTransactor) CommitVote(opts *bind.TransactOpts, _voter common.Address, commitment [32]byte, _numVotes uint32) (*types.Transaction, error) {
	return _QVote.contract.Transact(opts, "commitVote", _voter, commitment, _numVotes)
}

// CommitVote is a paid mutator transaction binding the contract method 0x03b8a89f.
//
// Solidity: function commitVote(address _voter, bytes32 commitment, uint32 _numVotes) returns()
func (_QVote *QVoteSession) CommitVote(_voter common.Address, commitment [32]byte, _numVotes uint32) (*types.Transaction, error) {
	return _QVote.Contract.CommitVote(&_QVote.TransactOpts, _voter, commitment, _numVotes)
}

// CommitVote is a paid mutator transaction binding the contract method 0x03b8a89f.
//
// Solidity: function commitVote(address _voter, bytes32 commitment, uint32 _numVotes) returns()
func (_QVote *QVoteTransactorSession) CommitVote(_voter common.Address, commitment [32]byte, _numVotes uint32) (*types.Transaction, error) {
	return _QVote.Contract.CommitVote(&_QVote.TransactOpts, _voter, commitment, _numVotes)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns(bool)
func (_QVote *QVoteTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QVote.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns(bool)
func (_QVote *QVoteSession) Destroy() (*types.Transaction, error) {
	return _QVote.Contract.Destroy(&_QVote.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns(bool)
func (_QVote *QVoteTransactorSession) Destroy() (*types.Transaction, error) {
	return _QVote.Contract.Destroy(&_QVote.TransactOpts)
}

// RevealVote is a paid mutator transaction binding the contract method 0x3932d6b7.
//
// Solidity: function revealVote(address _voter, uint256 _vote, bytes32 _salt) returns()
func (_QVote *QVoteTransactor) RevealVote(opts *bind.TransactOpts, _voter common.Address, _vote *big.Int, _salt [32]byte) (*types.Transaction, error) {
	return _QVote.contract.Transact(opts, "revealVote", _voter, _vote, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x3932d6b7.
//
// Solidity: function revealVote(address _voter, uint256 _vote, bytes32 _salt) returns()
func (_QVote *QVoteSession) RevealVote(_voter common.Address, _vote *big.Int, _salt [32]byte) (*types.Transaction, error) {
	return _QVote.Contract.RevealVote(&_QVote.TransactOpts, _voter, _vote, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x3932d6b7.
//
// Solidity: function revealVote(address _voter, uint256 _vote, bytes32 _salt) returns()
func (_QVote *QVoteTransactorSession) RevealVote(_voter common.Address, _vote *big.Int, _salt [32]byte) (*types.Transaction, error) {
	return _QVote.Contract.RevealVote(&_QVote.TransactOpts, _voter, _vote, _salt)
}

// ToggleContractActive is a paid mutator transaction binding the contract method 0x1385d24c.
//
// Solidity: function toggleContractActive() returns()
func (_QVote *QVoteTransactor) ToggleContractActive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QVote.contract.Transact(opts, "toggleContractActive")
}

// ToggleContractActive is a paid mutator transaction binding the contract method 0x1385d24c.
//
// Solidity: function toggleContractActive() returns()
func (_QVote *QVoteSession) ToggleContractActive() (*types.Transaction, error) {
	return _QVote.Contract.ToggleContractActive(&_QVote.TransactOpts)
}

// ToggleContractActive is a paid mutator transaction binding the contract method 0x1385d24c.
//
// Solidity: function toggleContractActive() returns()
func (_QVote *QVoteTransactorSession) ToggleContractActive() (*types.Transaction, error) {
	return _QVote.Contract.ToggleContractActive(&_QVote.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QVote *QVoteTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QVote.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QVote *QVoteSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QVote.Contract.TransferOwnership(&_QVote.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QVote *QVoteTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QVote.Contract.TransferOwnership(&_QVote.TransactOpts, newOwner)
}

// UploadFactor is a paid mutator transaction binding the contract method 0x8db372b7.
//
// Solidity: function uploadFactor(uint256 _factor, uint256 _userId) returns(bool)
func (_QVote *QVoteTransactor) UploadFactor(opts *bind.TransactOpts, _factor *big.Int, _userId *big.Int) (*types.Transaction, error) {
	return _QVote.contract.Transact(opts, "uploadFactor", _factor, _userId)
}

// UploadFactor is a paid mutator transaction binding the contract method 0x8db372b7.
//
// Solidity: function uploadFactor(uint256 _factor, uint256 _userId) returns(bool)
func (_QVote *QVoteSession) UploadFactor(_factor *big.Int, _userId *big.Int) (*types.Transaction, error) {
	return _QVote.Contract.UploadFactor(&_QVote.TransactOpts, _factor, _userId)
}

// UploadFactor is a paid mutator transaction binding the contract method 0x8db372b7.
//
// Solidity: function uploadFactor(uint256 _factor, uint256 _userId) returns(bool)
func (_QVote *QVoteTransactorSession) UploadFactor(_factor *big.Int, _userId *big.Int) (*types.Transaction, error) {
	return _QVote.Contract.UploadFactor(&_QVote.TransactOpts, _factor, _userId)
}

// UploadHash is a paid mutator transaction binding the contract method 0x8bb89e49.
//
// Solidity: function uploadHash(bytes32 _hash, bytes _signature, uint256 _userId) returns()
func (_QVote *QVoteTransactor) UploadHash(opts *bind.TransactOpts, _hash [32]byte, _signature []byte, _userId *big.Int) (*types.Transaction, error) {
	return _QVote.contract.Transact(opts, "uploadHash", _hash, _signature, _userId)
}

// UploadHash is a paid mutator transaction binding the contract method 0x8bb89e49.
//
// Solidity: function uploadHash(bytes32 _hash, bytes _signature, uint256 _userId) returns()
func (_QVote *QVoteSession) UploadHash(_hash [32]byte, _signature []byte, _userId *big.Int) (*types.Transaction, error) {
	return _QVote.Contract.UploadHash(&_QVote.TransactOpts, _hash, _signature, _userId)
}

// UploadHash is a paid mutator transaction binding the contract method 0x8bb89e49.
//
// Solidity: function uploadHash(bytes32 _hash, bytes _signature, uint256 _userId) returns()
func (_QVote *QVoteTransactorSession) UploadHash(_hash [32]byte, _signature []byte, _userId *big.Int) (*types.Transaction, error) {
	return _QVote.Contract.UploadHash(&_QVote.TransactOpts, _hash, _signature, _userId)
}

// VoteCommitPhase is a paid mutator transaction binding the contract method 0xfcd31320.
//
// Solidity: function voteCommitPhase() returns()
func (_QVote *QVoteTransactor) VoteCommitPhase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QVote.contract.Transact(opts, "voteCommitPhase")
}

// VoteCommitPhase is a paid mutator transaction binding the contract method 0xfcd31320.
//
// Solidity: function voteCommitPhase() returns()
func (_QVote *QVoteSession) VoteCommitPhase() (*types.Transaction, error) {
	return _QVote.Contract.VoteCommitPhase(&_QVote.TransactOpts)
}

// VoteCommitPhase is a paid mutator transaction binding the contract method 0xfcd31320.
//
// Solidity: function voteCommitPhase() returns()
func (_QVote *QVoteTransactorSession) VoteCommitPhase() (*types.Transaction, error) {
	return _QVote.Contract.VoteCommitPhase(&_QVote.TransactOpts)
}

// QVoteVoteCompletedIterator is returned from FilterVoteCompleted and is used to iterate over the raw logs and unpacked data for VoteCompleted events raised by the QVote contract.
type QVoteVoteCompletedIterator struct {
	Event *QVoteVoteCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QVoteVoteCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QVoteVoteCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QVoteVoteCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QVoteVoteCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QVoteVoteCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QVoteVoteCompleted represents a VoteCompleted event raised by the QVote contract.
type QVoteVoteCompleted struct {
	TopChoice *big.Int
	Candidate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoteCompleted is a free log retrieval operation binding the contract event 0xa88f280a1abb3c04cb45c068b9cd8d2fe5979cf6c6f86baa7164ae869e72e292.
//
// Solidity: event VoteCompleted(uint256 topChoice, address candidate)
func (_QVote *QVoteFilterer) FilterVoteCompleted(opts *bind.FilterOpts) (*QVoteVoteCompletedIterator, error) {

	logs, sub, err := _QVote.contract.FilterLogs(opts, "VoteCompleted")
	if err != nil {
		return nil, err
	}
	return &QVoteVoteCompletedIterator{contract: _QVote.contract, event: "VoteCompleted", logs: logs, sub: sub}, nil
}

// WatchVoteCompleted is a free log subscription operation binding the contract event 0xa88f280a1abb3c04cb45c068b9cd8d2fe5979cf6c6f86baa7164ae869e72e292.
//
// Solidity: event VoteCompleted(uint256 topChoice, address candidate)
func (_QVote *QVoteFilterer) WatchVoteCompleted(opts *bind.WatchOpts, sink chan<- *QVoteVoteCompleted) (event.Subscription, error) {

	logs, sub, err := _QVote.contract.WatchLogs(opts, "VoteCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QVoteVoteCompleted)
				if err := _QVote.contract.UnpackLog(event, "VoteCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCompleted is a log parse operation binding the contract event 0xa88f280a1abb3c04cb45c068b9cd8d2fe5979cf6c6f86baa7164ae869e72e292.
//
// Solidity: event VoteCompleted(uint256 topChoice, address candidate)
func (_QVote *QVoteFilterer) ParseVoteCompleted(log types.Log) (*QVoteVoteCompleted, error) {
	event := new(QVoteVoteCompleted)
	if err := _QVote.contract.UnpackLog(event, "VoteCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// QVoteVoteCreatedIterator is returned from FilterVoteCreated and is used to iterate over the raw logs and unpacked data for VoteCreated events raised by the QVote contract.
type QVoteVoteCreatedIterator struct {
	Event *QVoteVoteCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QVoteVoteCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QVoteVoteCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QVoteVoteCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QVoteVoteCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QVoteVoteCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QVoteVoteCreated represents a VoteCreated event raised by the QVote contract.
type QVoteVoteCreated struct {
	Creator     common.Address
	StartTime   *big.Int
	CloseTime   *big.Int
	Description [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoteCreated is a free log retrieval operation binding the contract event 0x5280092bcfbb26606763f7d314a01a5b9edb5de5f749abc4cb3f1c04067ca0eb.
//
// Solidity: event VoteCreated(address indexed _creator, uint256 _startTime, uint256 _closeTime, bytes32 _description)
func (_QVote *QVoteFilterer) FilterVoteCreated(opts *bind.FilterOpts, _creator []common.Address) (*QVoteVoteCreatedIterator, error) {

	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _QVote.contract.FilterLogs(opts, "VoteCreated", _creatorRule)
	if err != nil {
		return nil, err
	}
	return &QVoteVoteCreatedIterator{contract: _QVote.contract, event: "VoteCreated", logs: logs, sub: sub}, nil
}

// WatchVoteCreated is a free log subscription operation binding the contract event 0x5280092bcfbb26606763f7d314a01a5b9edb5de5f749abc4cb3f1c04067ca0eb.
//
// Solidity: event VoteCreated(address indexed _creator, uint256 _startTime, uint256 _closeTime, bytes32 _description)
func (_QVote *QVoteFilterer) WatchVoteCreated(opts *bind.WatchOpts, sink chan<- *QVoteVoteCreated, _creator []common.Address) (event.Subscription, error) {

	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _QVote.contract.WatchLogs(opts, "VoteCreated", _creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QVoteVoteCreated)
				if err := _QVote.contract.UnpackLog(event, "VoteCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCreated is a log parse operation binding the contract event 0x5280092bcfbb26606763f7d314a01a5b9edb5de5f749abc4cb3f1c04067ca0eb.
//
// Solidity: event VoteCreated(address indexed _creator, uint256 _startTime, uint256 _closeTime, bytes32 _description)
func (_QVote *QVoteFilterer) ParseVoteCreated(log types.Log) (*QVoteVoteCreated, error) {
	event := new(QVoteVoteCreated)
	if err := _QVote.contract.UnpackLog(event, "VoteCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// QVoteVoterApprovedIterator is returned from FilterVoterApproved and is used to iterate over the raw logs and unpacked data for VoterApproved events raised by the QVote contract.
type QVoteVoterApprovedIterator struct {
	Event *QVoteVoterApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QVoteVoterApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QVoteVoterApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QVoteVoterApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QVoteVoterApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QVoteVoterApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QVoteVoterApproved represents a VoterApproved event raised by the QVote contract.
type QVoteVoterApproved struct {
	Voter       common.Address
	Description [32]byte
	StartTime   *big.Int
	CloseTime   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoterApproved is a free log retrieval operation binding the contract event 0xd287efc8e32e1a20dd1f26bf23c7fcfdcd17404941f51e6afe527d87ed6289ad.
//
// Solidity: event VoterApproved(address indexed _voter, bytes32 _description, uint256 _startTime, uint256 _closeTime)
func (_QVote *QVoteFilterer) FilterVoterApproved(opts *bind.FilterOpts, _voter []common.Address) (*QVoteVoterApprovedIterator, error) {

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _QVote.contract.FilterLogs(opts, "VoterApproved", _voterRule)
	if err != nil {
		return nil, err
	}
	return &QVoteVoterApprovedIterator{contract: _QVote.contract, event: "VoterApproved", logs: logs, sub: sub}, nil
}

// WatchVoterApproved is a free log subscription operation binding the contract event 0xd287efc8e32e1a20dd1f26bf23c7fcfdcd17404941f51e6afe527d87ed6289ad.
//
// Solidity: event VoterApproved(address indexed _voter, bytes32 _description, uint256 _startTime, uint256 _closeTime)
func (_QVote *QVoteFilterer) WatchVoterApproved(opts *bind.WatchOpts, sink chan<- *QVoteVoterApproved, _voter []common.Address) (event.Subscription, error) {

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _QVote.contract.WatchLogs(opts, "VoterApproved", _voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QVoteVoterApproved)
				if err := _QVote.contract.UnpackLog(event, "VoterApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoterApproved is a log parse operation binding the contract event 0xd287efc8e32e1a20dd1f26bf23c7fcfdcd17404941f51e6afe527d87ed6289ad.
//
// Solidity: event VoterApproved(address indexed _voter, bytes32 _description, uint256 _startTime, uint256 _closeTime)
func (_QVote *QVoteFilterer) ParseVoterApproved(log types.Log) (*QVoteVoterApproved, error) {
	event := new(QVoteVoterApproved)
	if err := _QVote.contract.UnpackLog(event, "VoterApproved", log); err != nil {
		return nil, err
	}
	return event, nil
}
