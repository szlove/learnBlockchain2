package blockchain

import (
	"fmt"
	"log"
	"strings"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

const (
	MINING_DIFFYCALTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

type Blockchain struct {
	address         string
	chain           []*Block
	transactionPool []*Transaction
}

func NewBlockchain(address string) *Blockchain {
	bc := &Blockchain{address: address}
	genesis := &Block{}
	bc.CreateBlock(genesis.Hash(), 0)
	return bc
}

func (bc *Blockchain) CreateBlock(previousHash [32]byte, nonce int) *Block {
	b := NewBlock(previousHash, nonce, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) AddTransaction(sender, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
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

func (bc *Blockchain) ValidProof(previousHash [32]byte, nonce int, ts []*Transaction) bool {
	zeros := strings.Repeat("0", MINING_DIFFYCALTY)
	guessBlock := &Block{0, previousHash, nonce, ts}
	guessHashString := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashString[:MINING_DIFFYCALTY] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	lastBlock := bc.LastBlock()
	nonce := 0
	for !bc.ValidProof(lastBlock.Hash(), nonce, transactions) {
		nonce++
	}
	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.address, MINING_REWARD)
	lastBlock := bc.LastBlock()
	nonce := bc.ProofOfWork()
	bc.CreateBlock(lastBlock.Hash(), nonce)
	log.Println("action=mining status=success")
	return true
}

func (bc *Blockchain) Print() {
	for i, b := range bc.chain {
		r := strings.Repeat("=", 25)
		fmt.Printf("%s Block %d %s\n", r, i, r)
		b.Print()
	}
	fmt.Println(strings.Repeat("*", 50))
}
