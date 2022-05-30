package main

import (
	"fmt"

	"github.com/szlove/learnBlockchain2/blockchain"
)

func main() {
	address := "Slam"
	bc := blockchain.NewBlockchain(address)

	bc.AddTransaction("Slam", "Bee", 3.0)
	bc.Mining()

	bc.AddTransaction("Bee", "Slam", 4.5)
	bc.AddTransaction("Dog", "Slam", 400.0)
	bc.AddTransaction("Cat", "Slam", 800.0)
	bc.Mining()

	bc.AddTransaction("Slam", "Cat", 300.0)
	bc.AddTransaction("Dog", "Bee", 200.0)
	bc.Mining()

	bc.Print()

	fmt.Printf("Slam    %f\n", bc.CalculateTotalAmount("Slam"))
	fmt.Printf("Bee     %f\n", bc.CalculateTotalAmount("Bee"))
	fmt.Printf("Dog     %f\n", bc.CalculateTotalAmount("Dog"))
	fmt.Printf("Cat     %f\n", bc.CalculateTotalAmount("Cat"))
}
