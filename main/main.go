package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("hello")

	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/48709fc01e8245a6a61d4e2a9b1970bb")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey := ""

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf(err)
	}

	cAddress := common.HexToAddress("0x067D02EE461F1334F7aea27b52Ad687708D53D19")
	t, err := NewTodo(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// create a transaction signer
	tx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx.GasLimit = 5000000
	tx.GasPrice = gasPrice

	// add a todo task
	tran, err := t.AddTask(tx, "complete assignment")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tran.Hash())

	// get the entire taskslist
	pubKeyAddress := crypto.PubkeyToAddress(privateKey)
	tasks, err := t.GetTaskList(&bind.CallOpts{
		From: pubKeyAddress,
	})

	fmt.Println(tasks)

	// update the todo task data
	updateTran, err := t.UpdateTaskData(tx, big.NewInt(0), "play table tennis")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updateTran)
}
