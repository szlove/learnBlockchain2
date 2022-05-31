package main

import (
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
		wallet.B, 1.0)

	bc := blockchain.NewBlockchain(walletM.BlockchainAddress())

	bc.AddTransaction()
}
