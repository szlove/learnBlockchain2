package main

import (
	"github.com/szlove/learnBlockchain2/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain("my_blockchain_address")
	bc.AddTransaction("Slam", "Subin", 1.0)
	bc.Mining()
	bc.Print()
}
