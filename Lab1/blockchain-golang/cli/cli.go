package cli

import (
	// "fmt"
	// "encoding/json"
	"encoding/hex"
	"fmt"
	"golang-blockchain/blockchain"
	"golang-blockchain/database"
	"strconv"
	"strings"

	// "golang-blockchain/database"
	"os"

	"github.com/urfave/cli"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Run(client *mongo.Client) {
	var bc *blockchain.Blockchain
	app := cli.NewApp()
	app.Name = "Blockchain CLI"
	app.Usage = ""
	app.Version = "1.0.0"
	app.Copyright = "Group 6 since 2023"
	app.UsageText = "A simple blockchain CLI written by Golang"
	app.Commands = []cli.Command{
		{
			Name: "createblockchain",
			Action: func(c *cli.Context) error {
				bc = blockchain.NewBlockchain()
				database.AddBlockChain(bc, client)
				block := bc.Blocks[0]
				blockchain.PrintBlock(block)
				// fmt.Println("Create Blockchain successfully")
				return nil
			},
		},

		{
			Name: "addtrasaction",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "data",
				},
			},

			Action: func(c *cli.Context) error {
				data := c.String("data")
				if data == "" {
					return nil
				}
				split_data := strings.Split(data, ",")

				transactions := []*blockchain.Transaction{}
				for _, input := range split_data {
					transactions = append(transactions, &blockchain.Transaction{Data: []byte(input)})
				}
				var chains []blockchain.Blockchain = database.FindBlock(client)
				if len(chains) == 0{
					panic("Have you create blockchain first?")
				}
				chains[0].AddBlock(transactions)
				
				filter := bson.D{{"blocks.0.timestamp", chains[0].Blocks[0].Timestamp}}
				update := bson.D{{"$set", bson.D{{"blocks", chains[0].Blocks}}}}

				database.UpdateBlock(client, filter, update)
				return nil
			},
		},

		{
			Name: "updatetransaction",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "blocknumber",
				},
				cli.StringFlag{
					Name: "transnumber",
				},
				cli.StringFlag{
					Name: "data",
				},
			},
			Action: func(c *cli.Context) error {

				block_number := c.String("blocknumber")
				trans_number := c.String("transnumber")
				data := c.String("data")

				if block_number == "" || trans_number == "" {
					return nil
				}

				int_block_number, err := strconv.Atoi(block_number)
				if err != nil {
					panic(err)
				}

				int_trans_number, err := strconv.Atoi(trans_number)
				if err != nil {
					panic(err)
				}
				var chains []blockchain.Blockchain = database.FindBlock(client)
				if len(chains) == 0{
					panic("Have you create blockchain first?")
				}
				blockchain.UpdateTransactionData(&chains[0], int_block_number, int_trans_number, data)

				filter := bson.D{{"blocks.0.timestamp", chains[0].Blocks[0].Timestamp}}
				update := bson.D{{"$set", bson.D{{"blocks", chains[0].Blocks}}}}

				err = database.UpdateBlock(client, filter, update)
				if err != nil {
					panic(err)
				}

				fmt.Println("Update Successfully")
				return nil
			},
		},
		{
			Name: "printblockchain",
			Action: func(cli *cli.Context) error {
				var chains []blockchain.Blockchain = database.FindBlock(client)
				if len(chains) == 0{
					panic("Have you create blockchain first?")
				}
				for i, block := range chains[0].Blocks {
					fmt.Println(i)
					blockchain.PrintBlock(block)
				}
				return nil
			},
		},
		{
			Name: "printmerkletree",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "blocknumber",
				},
			},
			Action: func(c *cli.Context) error {
				block_number := c.String("blocknumber")
				var chains []blockchain.Blockchain = database.FindBlock(client)
				if len(chains) == 0{
					panic("Have you create blockchain first?")
				}
				int_block_number, err := strconv.Atoi(block_number)
				if err != nil {
					panic(err)
				}

				blockchain.PrintMerkleTree(chains[0].Blocks[int_block_number])
				return nil
			},
		},
		{
			Name: "blockvalidate",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "blocknumber",
				},
			},
			Action: func(c *cli.Context) error {
				block_number := c.String("blocknumber")
				var chains []blockchain.Blockchain = database.FindBlock(client)
				if len(chains) == 0{
					panic("Have you create blockchain first?")
				}
				int_block_number, err := strconv.Atoi(block_number)
				if err != nil {
					panic(err)
				}
				if hex.EncodeToString(chains[0].Blocks[int_block_number].CalculateMerkleRoot()) == hex.EncodeToString(chains[0].Blocks[int_block_number].MerkleRoot) {
					fmt.Println("Valid Block!")
				}else {
					fmt.Println("Invalid Block!")
				}
				return nil
			},
		},
	}

	app.Run(os.Args)

}