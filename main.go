package main

import (
	"github.com/szlove/learnBlockchain2/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain("Slam")

	bc.AddTransaction("Slam", "Subin", 1.0)
	bc.Mining()

	bc.AddTransaction("Slam", "Subin", 10.0)
	bc.AddTransaction("Subin", "Bee", 2.0)
	bc.Mining()

	bc.AddTransaction("Bee", "Slam", 1.0)
	bc.AddTransaction("Slam", "Subin", 0.5)
	bc.Mining()

	bc.Print()
}
