package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("hello")

	conn, err := ethclient.Dial("https://rinkeby.infura.io/v3/48709fc01e8245a6a61d4e2a9b1970bb")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contract, err := NewToken(common.HexToAddress("0x067D02EE461F1334F7aea27b52Ad687708D53D19"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	name, err := contract.Name(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("contract name:", name)

}
