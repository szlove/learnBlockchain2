package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	MINING_DIFFICULTY = 4
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

func init() {
	log.SetPrefix("Blockchain: ")
}

type Block struct {
	timestamp    int64
	previousHash [32]byte
	nonce        int
	transactions []*Transaction
}

func NewBlock(previousHash [32]byte, nonce int, transactions []*Transaction) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		previousHash: previousHash,
		nonce:        nonce,
		transactions: transactions,
	}
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Nonce        int            `json:"nonce"`
		Transactions []*Transaction `json:"transactions"`
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
	for _, t := range b.transactions {
		t.Print()
	}
}

type Blockchain struct {
	chain             []*Block
	transactionPool   []*Transaction
	blockchainAddress string
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	bc := &Blockchain{blockchainAddress: blockchainAddress}
	b := &Block{}
	bc.CreateBlock(b.Hash(), 0)
	return bc
}

func (bc *Blockchain) CreateBlock(previousHash [32]byte, nonce int) *Block {
	b := NewBlock(previousHash, nonce, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
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

func (bc *Blockchain) AddTransaction(sender, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		nt := NewTransaction(t.senderBlockchainAddress, t.recipientBlockchainAddress, t.value)
		transactions = append(transactions, nt)
	}
	return transactions
}

func (bc *Blockchain) ValidProof(
	previousHash [32]byte, nonce int, transactions []*Transaction, difficulty int,
) (hashString string, ok bool) {
	zeors := strings.Repeat("0", difficulty)
	guessBlock := &Block{0, previousHash, nonce, transactions}
	guessHashString := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashString, guessHashString[:difficulty] == zeors
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	prevHash := bc.LastBlock().Hash()
	hash := ""
	ok := false
	nonce := 0
	for !ok {
		hash, ok = bc.ValidProof(prevHash, nonce, transactions, MINING_DIFFICULTY)
		fmt.Printf("\r%s", hash)
		nonce++
	}
	fmt.Println()
	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	prevHash := bc.LastBlock().Hash()
	nonce := bc.ProofOfWork()
	bc.CreateBlock(prevHash, nonce)
	log.Println("action=mining status=success")
	return true
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
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address         %s\n", tx.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address      %s\n", tx.recipientBlockchainAddress)
	fmt.Printf(" value                             %f\n", tx.value)
}

func main() {
	bc := NewBlockchain("Slam")

	bc.AddTransaction(MINING_SENDER, "Slam", 100.0)
	bc.AddTransaction(MINING_SENDER, "Subin", 100.0)
	bc.AddTransaction(MINING_SENDER, "Bee", 100.0)
	bc.Mining()

	bc.AddTransaction("Slam", "Subin", 1.0)
	bc.AddTransaction("Subin", "Slam", 0.5)
	bc.AddTransaction("Slam", "Bee", 0.5)
	bc.AddTransaction("Slam", "Subin", 2.0)
	bc.Mining()

	bc.Print()
}
