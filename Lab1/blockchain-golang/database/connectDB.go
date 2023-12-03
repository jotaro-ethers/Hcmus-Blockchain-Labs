package database

import (
	"context"
	// "encoding/json"
	// "encoding/json"
	"fmt"
	"golang-blockchain/blockchain"
	"log"
	"os"

	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectDatabase()*mongo.Client{
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == ""{
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client,err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	
	if err != nil{
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")


	return client
}

func AddBlockChain(bc *blockchain.Blockchain,client *mongo.Client){

	collection := client.Database("bchain-lab1").Collection("Block")
	collection.InsertOne(context.TODO(),bc)

	
	// fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}

func FindBlock(client *mongo.Client) []blockchain.Blockchain{
	collection := client.Database("bchain-lab1").Collection("Block")
	cursor, err :=collection.Find(context.TODO(),bson.D{})

	if err != nil{
		panic(err)
	}

	var results []blockchain.Blockchain
	if err = cursor.All(context.TODO(), &results); err != nil {
		
		panic(err)
	}
	
	return results
	
}

func UpdateBlock(client *mongo.Client, filter interface{}, update interface{})error{
	collection := client.Database("bchain-lab1").Collection("Block")
	_,err:=collection.UpdateOne(context.TODO(),filter,update )
	return err
}