package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

type Block struct {
	timestamp    int64
	previousHash [32]byte
	nonce        int
	transactions []string
}

func NewBlock(previousHash [32]byte, nonce int) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		previousHash: previousHash,
		nonce:        nonce,
	}
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		PreviousHash [32]byte `json:"previous_hash"`
		Nonce        int      `json:"nonce"`
		Transactions []string `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		PreviousHash: b.previousHash,
		Nonce:        b.nonce,
		Transactions: b.transactions,
	})
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) Print() {
	fmt.Printf("timestamp          %d\n", b.timestamp)
	fmt.Printf("previousHash       %x\n", b.previousHash)
	fmt.Printf("nonce              %d\n", b.nonce)
	fmt.Printf("transactions       %s\n", b.transactions)
}

type Blockchain struct {
	chain           []*Block
	transactionPool []string
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	b := &Block{}
	bc.CreateBlock(b.Hash(), 0)
	return bc
}

func (bc *Blockchain) CreateBlock(previousHash [32]byte, nonce int) *Block {
	b := NewBlock(previousHash, nonce)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) LastBlock() *Block { return bc.chain[len(bc.chain)-1] }

func (bc *Blockchain) Print() {
	for i, b := range bc.chain {
		line := strings.Repeat("=", 25)
		fmt.Printf("%s Block %d %s\n", line, i, line)
		b.Print()
	}
	line := strings.Repeat("*", 25)
	fmt.Printf("%s %d Blocks %s\n", line, len(bc.chain), line)
}

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (tx *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchainAddress:    tx.senderBlockchainAddress,
		RecipientBlockchainAddress: tx.recipientBlockchainAddress,
		Value:                      tx.value,
	})
}

func (tx *Transaction) Print() {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf(" sender_blockchain_address         %s\n", tx.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address      %s\n", tx.recipientBlockchainAddress)
	fmt.Printf(" value                             %f\n", tx.value)
}

func main() {
	bc := NewBlockchain()
	for i := 0; i < 100; i++ {
		lb := bc.LastBlock()
		bc.CreateBlock(lb.Hash(), i)
	}
	bc.Print()
	tx := NewTransaction("Slam", "Bee", 1)
	tx.Print()
}
