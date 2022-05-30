package main

import (
	"fmt"

	"github.com/szlove/learnBlockchain2/wallet"
)

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKey())
	fmt.Println(w.PrivateKeyString())
	fmt.Println()
	fmt.Println(w.PublicKey())
	fmt.Println(w.PublicKeyString())
	fmt.Println()
	fmt.Println(w.BlockchainAddress())

	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "Bee", 1.0)
	fmt.Printf("Signature    %s\n", t.GenerateSignature())
}
