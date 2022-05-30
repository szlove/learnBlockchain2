package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func NewWallet() *Wallet {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return &Wallet{privateKey: privateKey, publicKey: &privateKey.PublicKey}
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey { return w.privateKey }
func (w *Wallet) PrivateKeyString() string      { return fmt.Sprintf("%x", w.privateKey.D.Bytes()) }
func (w *Wallet) PublicKey() *ecdsa.PublicKey   { return w.publicKey }
func (w *Wallet) PublicKeyString() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}
