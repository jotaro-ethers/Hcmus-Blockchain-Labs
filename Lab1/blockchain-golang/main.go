package main

import (
	// "fmt"

	"golang-blockchain/cli"
	"golang-blockchain/database"
)

func main() {
	client := database.ConnectDatabase()
	
	cli.Run(client)
}

