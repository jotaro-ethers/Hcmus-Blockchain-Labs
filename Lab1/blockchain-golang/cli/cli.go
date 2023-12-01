package cli

import (
	"os"
	"github.com/urfave/cli"
	"fmt"
	"golang-blockchain/blockchain"

)

func Run() {
	var bc *blockchain.Blockchain
	app := cli.NewApp()
	app.Name = "Blockchain CLI"
	app.Usage = ""
	app.Version = "1.0.0"
	app.Copyright="Group 6 since 2023"
	app.UsageText = "A simple blockchain CLI written by Golang"
	app.Commands = []cli.Command{
		{
			Name:"createblockchain" ,
			Action: func (c *cli.Context )error {
				bc = blockchain.NewBlockchain()
				blockchain.PrintBlock(bc.Blocks[0])
				fmt.Println("Create Blockchain successfully")
				return nil
			},
		},
	
	}
	
	app.Run(os.Args)

}