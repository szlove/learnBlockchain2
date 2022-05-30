package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/szlove/learnBlockchain2/util"
)

const (
	MINING_SENDER     string  = "THE BLOCKCHAIN"
	MINING_REWARD     float32 = 1.0
	MINING_DIFFICALTY int     = 3
)

type Blockchain struct {
	address         string
	chain           []*Block
	transactionPool []*Transaction
}

func NewBlockchain(address string) *Blockchain {
	bc := &Blockchain{}
	bc.address = address
	nilBlock := &Block{}
	bc.CreateBlock(nilBlock.Hash(), 0)
	return bc
}

func (bc *Blockchain) CreateBlock(previousHash [32]byte, nonce int) *Block {
	b := NewBlock(previousHash, nonce, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) AddTransaction(sender, recipient string, value float32,
	senderPublicKey *ecdsa.PublicKey, s *util.Signature) (ok bool) {
	t := NewTransaction(sender, recipient, value)
	if sender == MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}
	if bc.VerifyTransactionSignature(senderPublicKey, s, t) {
		if bc.CalculateTotalAmount(sender) < value {
			log.Println("ERROR: Not enough valance in a wallet")
			return false
		}
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}
	log.Println("ERROR: Verify transaction")
	return false
}

func (bc *Blockchain) VerifyTransactionSignature(
	senderPublicKey *ecdsa.PublicKey, s *util.Signature, t *Transaction) bool {
	m, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	h := sha256.Sum256(m)
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	var transactions []*Transaction
	for _, t := range bc.transactionPool {
		transactions = append(transactions, t)
	}
	return transactions
}

func (bc *Blockchain) LastBlock() *Block { return bc.chain[len(bc.chain)-1] }

func (bc *Blockchain) ValidProof(previousHash [32]byte, nonce int, transactions []*Transaction) bool {
	zeros := strings.Repeat("0", MINING_DIFFICALTY)
	guessBlock := &Block{0, previousHash, nonce, transactions}
	guessHashString := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashString[:MINING_DIFFICALTY] == zeros
}

func (bc *Blockchain) ProofOfWork(previousHash [32]byte) int {
	transactions := bc.CopyTransactionPool()
	nonce := 0
	for !bc.ValidProof(previousHash, nonce, transactions) {
		nonce++
	}
	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.address, MINING_REWARD, nil, nil)
	lastBlock := bc.LastBlock()
	nonce := bc.ProofOfWork(lastBlock.Hash())
	bc.CreateBlock(lastBlock.Hash(), nonce)
	return true
}

func (bc *Blockchain) Print() {
	for i, b := range bc.chain {
		r := strings.Repeat("=", 25)
		fmt.Printf("%s Block%d %s\n", r, i, r)
		b.Print()
	}
	fmt.Println(strings.Repeat("*", 50))
}

func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			switch blockchainAddress {
			case t.recipientBlockchainAddress:
				totalAmount += t.value
			case t.senderBlockchainAddress:
				totalAmount -= t.value
			}
		}
	}
	return totalAmount
}
