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

func (b *Block) String() {
	fmt.Printf("timestamp          %d\n", b.timestamp)
	fmt.Printf("previousHash       %x\n", b.previousHash)
	fmt.Printf("nonce              %d\n", b.nonce)
	fmt.Printf("transactions       %s\n", b.transactions)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64
		PreviousHash [32]byte
		Nonce        int
		Transaction  []string
	}{
		Timestamp:    b.timestamp,
		PreviousHash: b.previousHash,
		Nonce:        b.nonce,
		Transaction:  b.transactions,
	})
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		previousHash: previousHash,
		nonce:        nonce,
	}
}

type Blockchain struct {
	chain           []*Block
	transactionPool []string
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		r := strings.Repeat("=", 25)
		fmt.Printf("%s Chain %d %s\n", r, i, r)
		block.String()
	}
	fmt.Println(strings.Repeat("*", 50))
}

func (bc *Blockchain) LastBlock() *Block { return bc.chain[len(bc.chain)-1] }

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := &Blockchain{}
	bc.CreateBlock(0, b.Hash())
	return bc
}

func main() {
	bc := NewBlockchain()
	prevHash := bc.LastBlock().Hash()
	bc.CreateBlock(5, prevHash)
	prevHash = bc.LastBlock().Hash()
	bc.CreateBlock(10, prevHash)
	prevHash = bc.LastBlock().Hash()
	bc.CreateBlock(13, prevHash)
	bc.Print()
}
