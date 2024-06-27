package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    TransactionStorage "test.com/contracts"// for demo
)

func main() {
    client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/NJQIqo8Cga6zvoOrGw0UsXBJWE01K_hF")
    if err != nil {
        log.Fatal(err)
    }

    address := common.HexToAddress("0x75fA1F8f71A48AB5802Bc634f0eE7D0Cc76De8c1")
    instance, err := TransactionStorage.NewTransactionStorage(address, client)
    if err != nil {
        log.Fatal(err)
    }

    transactions, err := instance.GetTransactions(nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(transactions)
}