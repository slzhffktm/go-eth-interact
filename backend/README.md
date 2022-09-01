# Running The Go Program to Interact with The Smart Contract

This is following the [goethereumbook.org Smart Contracts Section](https://goethereumbook.org/en/smart-contracts/) and [Creating a simple Ethereum Smart Contract in Golang](https://towardsdev.com/creating-a-simple-ethereum-smart-contract-in-golang-138b9439f64e).

## Getting Started

1. Make sure you have `go` installed. This tutorial is using `go 1.18`
2. Copy the `.env.example` file to `.env` and fill in the values, 
   where the `CONTRACT_ADDRESS` is the address from the `contracts/.env`,
   and the `OWNER_ADDRESS` is the address that was used to deploy the smart contract
3. Go to `main.go` and change the message at line 39 with your name 
  ```
    resp2, err := instance.GreetOwner(&bind.CallOpts{
		From: common.HexToAddress(os.Getenv("OWNER_ADDRESS")),
	}, "<your name here>")
  ```
4. Run `make start` to start the program
5. You should see the response from the smart contract being logged
