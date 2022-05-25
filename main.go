package main

import (
	"github.com/szlove/learnBlockchain2/blockchain"
)

func main() {
	myBlockchainAddress := "my_blockchain_address"
	bc := blockchain.NewBlockchain(myBlockchainAddress)

	bc.AddTransaction("Slam", "Subin", 1.0)
	bc.AddTransaction("Subin", "Bee", 0.5)
	bc.AddTransaction("Bee", "Slam", 0.25)
	bc.Mining()

	bc.AddTransaction("Slam", "Bee", 10.0)
	bc.AddTransaction("Bee", "Suzi", 3.0)
	bc.AddTransaction("Suzi", "Subin", 1.55)
	bc.Mining()

	bc.AddTransaction("Suzi", "Bee", 1.0)
	bc.AddTransaction("Bee", "Subin", 1.0)
	bc.AddTransaction("Subin", "Slam", 3.3)
	bc.Mining()

	bc.Print()
}
