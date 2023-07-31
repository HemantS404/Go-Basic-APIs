package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "enter connection string"
const dbName = "netflix"
const watchlist = "watchlist"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo Connection Success")
	collection = client.Database(dbName).Collection(watchlist)
	fmt.Println("Collection reference ready")
}

func GetCollection() *mongo.Collection {
	return collection
}
