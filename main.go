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

func NewBlock(nonce int, previousHash [32]byte) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		previousHash: previousHash,
		nonce:        nonce,
	}
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64
		PreviousHash [32]byte
		Nonce        int
		Transactions []string
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

func (b *Block) String() {
	fmt.Printf("timestamp               %d\n", b.timestamp)
	fmt.Printf("previousHash            %x\n", b.previousHash)
	fmt.Printf("nonce                   %d\n", b.nonce)
	fmt.Printf("transactions            %s\n", b.transactions)
}

type Blockchain struct {
	chain           []*Block
	transactionPool []string
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	b := &Block{}
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) LastBlock() *Block { return bc.chain[len(bc.chain)-1] }

func (bc *Blockchain) Print() {
	for i, b := range bc.chain {
		r := strings.Repeat("=", 25)
		fmt.Printf("%s Block %d %s\n", r, i, r)
		b.String()
	}
	fmt.Println(strings.Repeat("*", 50))
}

func main() {
	bc := NewBlockchain()
	prevHash := bc.LastBlock().Hash()
	bc.CreateBlock(10, prevHash)
	bc.Print()
}
