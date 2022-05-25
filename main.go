package main

import (
	"fmt"

	"github.com/szlove/learnBlockchain2/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain("Slam")

	bc.AddTransaction("Slam", "Bee", 1.0)
	bc.AddTransaction("Subin", "Bee", 1.0)
	bc.Mining()

	bc.AddTransaction("Suzi", "Slam", 3.0)
	bc.AddTransaction("Bee", "Slam", 5.0)
	bc.Mining()

	bc.Print()

	fmt.Printf("Slam    %f\n", bc.CalculateTotalAmount("Slam"))
	fmt.Printf("Bee     %f\n", bc.CalculateTotalAmount("Bee"))
	fmt.Printf("Subin   %f\n", bc.CalculateTotalAmount("Subin"))
	fmt.Printf("Suzi    %f\n", bc.CalculateTotalAmount("Suzi"))
}
