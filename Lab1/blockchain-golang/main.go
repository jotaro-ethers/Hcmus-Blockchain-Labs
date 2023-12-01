package main

import (
	"golang-blockchain/cli"
)

func main() {
	// bc := blockchain.NewBlockchain()

	// // Tạo các giao dịch
	// transactions1 := []*blockchain.Transaction{&blockchain.Transaction{Data: []byte("Send 1 BTC to Vo Le Hoai")}, &blockchain.Transaction{Data: []byte("Send 10000 BTC to Nguyen Vu Khoi")},&blockchain.Transaction{Data: []byte("Send 1 ETH to Nguyen Vu Khoi")}}
	// transactions2 := []*blockchain.Transaction{&blockchain.Transaction{Data: []byte("Send 2 more BTC to Pham Dang Quan")}}

	// // Thêm các khối vào chuỗi khối
	// bc.AddBlock(transactions1)
	// bc.AddBlock(transactions2)

	// for _, block := range bc.Blocks {
    //     fmt.Printf("Block Number: %d\n", block.BlockNumber)
	// 	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	// 	fmt.Println("Transactions:")
	// 	for _, tx := range block.Transactions {
	// 		fmt.Printf("  - %s\n", string(tx.Data))
	// 	}
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Printf("Timestamp: %d\n", block.Timestamp)
	// 	fmt.Printf("MerkleRoot: %x\n", block.MerkleRoot)
	// 	fmt.Println()
	// }

    // fmt.Printf("Total Block: %d\n", len(bc.Blocks))
    // blockchain.PrintMerkleTree(bc.Blocks[1])
	// fmt.Println("Update transaction 0 block 1")
	// blockchain.UpdateTransactionData(bc, 1, 0, "Send 2 BTC to Pham Dang Quan")
	// blockchain.PrintMerkleTree(bc.Blocks[1])
	cli.Run()

}
