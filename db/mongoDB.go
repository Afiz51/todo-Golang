package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Users *mongo.Collection

func ConnectToMongoDB(database string) {
	// MongoDB connection URI
	uri := "mongodb://localhost:27017"

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	Client = client

	Users = client.Database(database).Collection("users")

}

func PrintCollection(databaseName, collectionName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Client.Database(databaseName).Collection(collectionName)

	// Define a filter to retrieve all documents
	filter := bson.D{}

	// Perform a find operation to retrieve documents matching the filter
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatalf("Error finding documents in collection: %v", err)
	}
	defer cursor.Close(ctx)

	// Iterate through the cursor and print each document
	if cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatalf("Error decoding document: %v", err)
		}

		// Print the JSON representation of the document
		if jsonData, err := bson.MarshalExtJSON(result, false, false); err != nil {
			log.Fatalf("Error marshaling document to JSON: %v", err)
		} else {
			log.Println(string(jsonData))
		}
	}
	if err := cursor.Err(); err != nil {
		log.Fatalf("Error iterating cursor: %v", err)
	}
}
