# Running The Go Program to Interact with The Smart Contract

This is following the [goethereumbook.org Smart Contract Write Section](https://goethereumbook.org/en/smart-contract-write/) and [Creating a simple Ethereum Smart Contract in Golang](https://towardsdev.com/creating-a-simple-ethereum-smart-contract-in-golang-138b9439f64e).


## Getting Started

1. Check the `contracts/helloworld.go` file. This file is the SDK for interacting with the smart contract.
   This file is generated automatically from the smart contract by running the `make generate-go-api` from the root directory.
2. Make sure you have `go` installed. This tutorial is using `go 1.18`
3. Copy the `.env.example` file to `.env` and fill in the values, 
   where the `CONTRACT_ADDRESS` and `PRIVATE_KEY` are from the `contracts/.env`,
4. Go to `main.go` and change the message at line 71 with your name 
   ```
     tx, err := instance.UpdateOwnerName(auth, "<your name here>")
   ```
5. Run `make start` to start the program
6. You should see the responses from the smart contract being logged
