package wallet

import "crypto/ecdsa"

type Transaction struct {
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(
	senderPrivateKey *ecdsa.PrivateKey,
	senderPublicKey *ecdsa.PublicKey,
	senderBlockchainAddress string,
	recipientBlockchainAddress string,
	value float32,
) *Transaction {
	return &Transaction{
		senderPrivateKey,
		senderPublicKey,
		senderBlockchainAddress,
		recipientBlockchainAddress,
		value,
	}
}
