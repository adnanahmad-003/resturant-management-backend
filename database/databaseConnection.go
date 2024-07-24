package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func DBinstance() *mongo.Client {
	// Create a direct connection to a host. The driver will send all requests
	// to that host and will not automatically discover other hosts in the
	// deployment.
	// 	context.WithTimeout creates a context that will automatically be canceled after 10 seconds. This is useful to avoid waiting indefinitely if something goes wrong while connecting to MongoDB.
	// defer cancel() ensures that the context's cancel function is called to release resources once the DBinstance function returns
	// mongo.Connect is used to create a new MongoDB client and initiate a connection to the server using the provided context and client options. If there is an error during the connection, log.Fatal(err) will log the error and exit the program

	clientOpts := options.Client().ApplyURI(
		"mongodb://localhost:27017/?connect=direct")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongoDb")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = Client.Database("resturant").Collection(collectionName)
	return collection
}
