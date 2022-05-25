package main

import (
	"fmt"
	"log"

	"github.com/szlove/learnBlockchain2/blockchain"
	"github.com/szlove/learnBlockchain2/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(),
		walletB.BlockchainAddress(), 1.0)

	bc := blockchain.NewBlockchain(walletM.BlockchainAddress())
	isAdded := bc.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0,
		walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added?", isAdded)

	bc.Mining()
	bc.Print()

	fmt.Printf("A    %f\n", bc.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B    %f\n", bc.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M    %f\n", bc.CalculateTotalAmount(walletM.BlockchainAddress()))
}
