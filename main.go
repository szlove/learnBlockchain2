package main

import (
	"fmt"

	"github.com/szlove/learnBlockchain2/blockchain"
)

func main() {
	myBlockchainAddress := "my_blockchain_address"
	bc := blockchain.NewBlockchain(myBlockchainAddress)

	bc.AddTransaction("Slam", "Suzi", 3.0)
	bc.AddTransaction("Subin", "Suzi", 1.5)
	bc.Mining()

	bc.AddTransaction("Suzi", "Slam", 2.2)
	bc.AddTransaction("Subin", "Suzi", 3.3)
	bc.Mining()

	bc.Print()

	fmt.Printf("Slam    %f\n", bc.CalculateTotalAmount("Slam"))
	fmt.Printf("Suzi    %f\n", bc.CalculateTotalAmount("Suzi"))
	fmt.Printf("Subin    %f\n", bc.CalculateTotalAmount("Subin"))
}
