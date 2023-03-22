package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	uri := os.Getenv("MONGO_URI")
	log.Printf("Mongo uri %v", uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatalf("Failed to connect to db %v", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatalf("Couldnt establish connection %v", err)
	}
	defer cancel()

	log.Println("Connected to Database")

	return client
}


var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("demo").Collection(collectionName)
	return collection
}
