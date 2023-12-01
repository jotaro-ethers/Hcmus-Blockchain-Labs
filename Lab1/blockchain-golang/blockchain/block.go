// Block trong block.go
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
	"errors"
	// "fmt"
)

type Block struct {
	BlockNumber   int 
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	MerkleRoot    []byte
}

type Transaction struct {
	Data []byte
}

type Blockchain struct {
	Blocks []*Block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	var transactionsData [][]byte
	for _, tx := range b.Transactions {
		transactionsData = append(transactionsData, tx.Data)
	}

	// Nối các mảng byte lại với nhau
	headers := bytes.Join([][]byte{b.PrevBlockHash, bytes.Join(transactionsData, []byte{}), timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(blockNumber int, transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{blockNumber, time.Now().Unix(), transactions, prevBlockHash, []byte{}, []byte{}}
	block.SetHash()
	block.MerkleRoot = block.CalculateMerkleRoot()
	return block
}

func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlockNumber := prevBlock.BlockNumber + 1
	newBlock := NewBlock(newBlockNumber, transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (b *Block) CalculateMerkleRoot() []byte {
	var transactionsData [][]byte
	for _, tx := range b.Transactions {
		transactionsData = append(transactionsData, tx.Data)
	}
	merkleTree := NewMerkleTree(transactionsData)
	return merkleTree.RootNode.Data
}

func NewGenesisBlock(initData string) *Block {
	initTransaction := &Transaction{Data: []byte(initData)}
	return NewBlock(0, []*Transaction{initTransaction}, []byte{})
}

func NewBlockchain() *Blockchain {
	initData := "Init Genesis Block"
	genesisBlock := NewGenesisBlock(initData)
	return &Blockchain{[]*Block{genesisBlock}}
}

func UpdateTransactionData(blockchain *Blockchain, blockNumber, transactionIndex int, newTransactionData string) error {
	if blockNumber < 0 || blockNumber >= len(blockchain.Blocks) {
		return errors.New("Invalid block number")
	}
	if transactionIndex < 0 || transactionIndex >= len(blockchain.Blocks[blockNumber].Transactions) {
		return errors.New("Invalid transaction index")
	}

	blockCopy := *blockchain.Blocks[blockNumber]
	blockCopy.Transactions[transactionIndex].Data = []byte(newTransactionData)
	blockchain.Blocks[blockNumber] = &blockCopy
	return nil
}
// func printMai(block *Block){
// 	fmt.Printf("Block Number: %d\n", block.BlockNumber)
// 	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
// 	fmt.Println("Transactions:")
// 	for _, tx := range block.Transactions {
// 		fmt.Printf("  - %s\n", string(tx.Data))
// 	}
// 	fmt.Printf("Hash: %x\n", block.Hash)
// 	fmt.Printf("Timestamp: %d\n", block.Timestamp)
// 	fmt.Printf("MerkleRoot: %x\n", block.MerkleRoot)
// }
	