.PHONY:geneth

mkdir:
	$(shell mkdir -p ./VoteSrc/ethbridge/go/TaskCore)
	$(shell mkdir -p ./VoteSrc/ethbridge/go/QVote)

geneth:
	solc --optimize --bin -o ./build ./contracts/TaskCore.sol --overwrite
	solc --abi -o ./build ./contracts/TaskCore.sol --overwrite
	./abigeneth.exe --bin=./build/TaskCore.bin --abi=./build/TaskCore.abi --pkg=TaskCore --out=./eth/TaskCore/TaskCore.go
	solc --optimize --bin -o ./build ./contracts/QVote.sol --overwrite
	solc --abi -o ./build ./contracts/QVote.sol --overwrite
	./abigeneth.exe --bin=./build/QVote.bin --abi=./build/QVote.abi --pkg=QVote --out=./eth/QVote/QVote.go


# ganache-cli -m "much repair shock carbon improve miss forget sock include bullet interest solution"
