package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

type Block struct {
	timestamp    int64
	previousHash string
	nonce        int
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		previousHash: previousHash,
		nonce:        nonce,
	}
}

func main() {
	b := NewBlock(0, "genesis")
	fmt.Println(b)
}
