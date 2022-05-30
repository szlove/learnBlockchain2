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
}
