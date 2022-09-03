# Running The Go Program to Interact with The Smart Contract

This is following the [goethereumbook.org Smart Contract Write Section](https://goethereumbook.org/en/smart-contract-write/) and [Creating a simple Ethereum Smart Contract in Golang](https://towardsdev.com/creating-a-simple-ethereum-smart-contract-in-golang-138b9439f64e).

## Getting Started

1. Make sure you have `go` installed. This tutorial is using `go 1.18`
2. Copy the `.env.example` file to `.env` and fill in the values, 
   where the `CONTRACT_ADDRESS` and `PRIVATE_KEY` are from the `contracts/.env`,
3. Go to `main.go` and change the message at line 71 with your name 
  ```
    tx, err := instance.UpdateOwnerName(auth, "<your name here>")
  ```
4. Run `make start` to start the program
5. You should see the responses from the smart contract being logged
