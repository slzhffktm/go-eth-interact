generate-go-api:
	solc @openzeppelin/=contracts/node_modules/@openzeppelin/ contracts/contracts/HelloWorld.sol --optimize --overwrite --abi contracts/contracts/HelloWorld.sol -o temp
	abigen --abi=./temp/HelloWorld.abi --pkg=helloworld --out=./backend/contracts/helloworld.go
