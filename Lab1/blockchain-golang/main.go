package main

import (
	"golang-blockchain/blockchain"
	"fmt"
	"bufio"
	"os"
	"encoding/hex"

)

func main() {
	var bc *blockchain.Blockchain
	bc = &blockchain.Blockchain{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("##############################################")
		fmt.Println("1. Create Blockchain \n2. Add Block \n3. Print Blockchain \n4. Update Transaction \n5. Print Merkle Tree \n6. Validate Block \n7. Exit")
		
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			bc = blockchain.NewBlockchain()
			blockchain.PrintBlock(bc.Blocks[0])
		case 2:
			if len(bc.Blocks) == 0 {
				fmt.Println("Create Blockchain first")
				break
			}
			fmt.Println("New Transaction, One per line, Enter to stop")
			transactions := []*blockchain.Transaction{}
			for {
				scanner.Scan()
				input := scanner.Text()
				if len(input) > 0 {
					transactions = append(transactions, &blockchain.Transaction{Data: []byte(input)})
				}else {
					break
				}
			}
			bc.AddBlock(transactions)
		case 3:
			if len(bc.Blocks) == 0 {
				fmt.Println("Create Blockchain first")
				break
			}
			for _, block := range bc.Blocks {
				blockchain.PrintBlock(block)
			}
		case 4:
			if len(bc.Blocks) == 0 {
				fmt.Println("Create Blockchain first")
				break
			}
			fmt.Println("Enter Block Number")
			var blockNumber int
			fmt.Scan(&blockNumber)
			fmt.Println("Enter Transaction Number")
			var transactionNumber int
			fmt.Scan(&transactionNumber)
			fmt.Println("Enter New Transaction Data")
			scanner.Scan()
			input := scanner.Text()
			blockchain.UpdateTransactionData(bc, blockNumber, transactionNumber, input)
		case 5:
			if len(bc.Blocks) == 0 {
				fmt.Println("Create Blockchain first")
				break
			}
			fmt.Println("Enter Block Number")
			var blockNumber int
			fmt.Scan(&blockNumber)
			blockchain.PrintMerkleTree(bc.Blocks[blockNumber])
		case 6:
			if len(bc.Blocks) == 0 {
				fmt.Println("Create Blockchain first")
				break
			}
			fmt.Println("Enter Block Number")
			var blockNumber int
			fmt.Scan(&blockNumber)
			block := bc.Blocks[blockNumber]
			merkleroot := hex.EncodeToString(block)
			fmt.Println(merkleroot)
			// if block.CalculateMerkleRoot() != bc.Blocks[blockNumber].MerkleRoot {
			// 	fmt.Println("Invalid Block")
			// }else {
			// 	fmt.Println("Valid Block")
			// }			
		}
		

	}

}

