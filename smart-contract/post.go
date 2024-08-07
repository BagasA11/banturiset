package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	TransactionStorage "test.com/contracts" // for demo
)

func post() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/NJQIqo8Cga6zvoOrGw0UsXBJWE01K_hF")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("f368b29aee79a42446ee9878c0f5bc057bc51ed5986985f192cfc13d5e55f0b5")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	println(nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	println(gasPrice)
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x75fA1F8f71A48AB5802Bc634f0eE7D0Cc76De8c1")
	instance, err := TransactionStorage.NewTransactionStorage(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// key := [32]byte{}
	// value := [32]byte{}
	// copy(key[:], []byte("foo"))
	// copy(value[:], []byte("bar"))

	from := "Joko"
	to := "Budi"
	amount := big.NewInt(int64(999))
	tx, err := instance.AddTransaction(auth, from, to, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	// result, err := instance.Items(nil, key)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// println("Hello, World!")
	// fmt.Println(string(result[:])) // "bar"
}
