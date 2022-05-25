package main

import (
	"fmt"
	"log"

	"github.com/szlove/learnBlockchain2/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyString())
	fmt.Println()
	fmt.Println(w.PublicKeyString())
	fmt.Println()
	fmt.Println(w.BlockchainAddress())
}
