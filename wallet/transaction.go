package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"

	"github.com/szlove/learnBlockchain2/util"
)

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

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchainAddress:    t.senderBlockchainAddress,
		RecipientBlockchainAddress: t.recipientBlockchainAddress,
		Value:                      t.value,
	})
}

func (t *Transaction) GenerateSignature() *util.Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256(m)
	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	return &util.Signature{r, s}
}
