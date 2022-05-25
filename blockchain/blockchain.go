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
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
	MINING_DIFFICULTY = 3
)

type Blockchain struct {
	address         string
	chain           []*Block
	transactionPool []*Transaction
}

func NewBlockchain(myBlockchainAddress string) *Blockchain {
	bc := &Blockchain{}
	bc.address = myBlockchainAddress
	genesisBlock := &Block{}
	bc.CreateBlock(genesisBlock.Hash(), 0)
	return bc
}

func (bc *Blockchain) CreateBlock(previousHash [32]byte, nonce int) *Block {
	b := NewBlock(previousHash, nonce, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) VerifyTransactionSignature(
	senderPublicKey *ecdsa.PublicKey, s *util.Signature, t *Transaction,
) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256(m)
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

func (bc *Blockchain) AddTransaction(
	sender, recipient string, value float32,
	senderPublicKey *ecdsa.PublicKey, s *util.Signature,
) bool {
	t := NewTransaction(sender, recipient, value)
	switch {
	case sender == MINING_SENDER, bc.VerifyTransactionSignature(senderPublicKey, s, t):
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	default:
		log.Println("ERROR: verify transaction")
		return false
	}
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	var transactions []*Transaction
	for _, t := range bc.transactionPool {
		transaction := NewTransaction(t.senderBlockchainAddress, t.recipientBlockchainAddress, t.value)
		transactions = append(transactions, transaction)
	}
	return transactions
}

func (bc *Blockchain) LastBlock() *Block { return bc.chain[len(bc.chain)-1] }

func (bc *Blockchain) ValidProof(previousHash [32]byte, nonce int, transactions []*Transaction) bool {
	zeros := strings.Repeat("0", MINING_DIFFICULTY)
	guessBlock := &Block{0, previousHash, nonce, transactions}
	guessHashString := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashString[:MINING_DIFFICULTY] == zeros
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
	log.Println("action=mining status=success")
	return true
}

func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			switch blockchainAddress {
			case t.senderBlockchainAddress:
				totalAmount -= t.value
				break
			case t.recipientBlockchainAddress:
				totalAmount += t.value
			}
		}
	}
	return totalAmount
}

func (bc *Blockchain) Print() {
	for i, b := range bc.chain {
		r := strings.Repeat("=", 25)
		fmt.Printf("%s Block %d %s\n", r, i, r)
		b.Print()
	}
	fmt.Println(strings.Repeat("*", 50))
}
