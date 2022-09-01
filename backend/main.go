package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	helloworld "github.com/slzhffktm/go-eth-interact/contracts"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("API_URL"))
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := helloworld.NewHelloworld(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// No owner address needed for this call.
	resp, err := instance.Hello(nil)
	if err != nil {
		log.Fatal(fmt.Errorf("Hello error: %w", err))
	}

	fmt.Printf("The respond from Hello() is: %s\n", resp)

	// Only the owner can call the next function.
	// Thus we need to specify the owner address.
	resp2, err := instance.GreetOwner(&bind.CallOpts{
		From: common.HexToAddress(os.Getenv("OWNER_ADDRESS")),
	}, "Owner")
	if err != nil {
		log.Fatal(fmt.Errorf("GreetOwner error: %w", err))
	}

	fmt.Printf("The respond from GreetOwner() is: %s\n", resp2)
}
