package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseInstance() *mongo.Client {
	mongoDB := "mongodb://localhost:27017"
	fmt.Print(mongoDB)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDB))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nConnected to mongodb")

	return client
}


var Client *mongo.Client = DatabaseInstance()

func openCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection = *mongo.Collection = client.Database("gorestaurant").Collection(collectionName)
	return collection
}