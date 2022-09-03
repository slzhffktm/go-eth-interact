package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	helloworld "github.com/slzhffktm/go-eth-interact/contracts"
)

func main() {
	ctx := context.Background()

	client, err := ethclient.DialContext(ctx, os.Getenv("API_URL"))
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := helloworld.NewHelloworld(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// No owner address needed for this call.
	resp, err := instance.Hello(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		log.Fatal(fmt.Errorf("Hello error: %w", err))
	}
	fmt.Printf("The respond from Hello() is: %s\n", resp)

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Owner needs to sign the transaction for calling this function.
	// Thus auth is needed.
	tx, err := instance.UpdateOwnerName(auth, "Owner from Go")
	if err != nil {
		log.Fatal(fmt.Errorf("UpdateOwnerName error: %w", err))
	}
	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())

	// Only the owner can call the next function.
	// Thus we need to specify the from address.
	resp2, err := instance.GreetOwner(&bind.CallOpts{
		From:    fromAddress,
		Context: ctx,
	})
	if err != nil {
		log.Fatal(fmt.Errorf("GreetOwner error: %w", err))
	}
	fmt.Printf("The respond from GreetOwner() is: %s\n", resp2)
}
