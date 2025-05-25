package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupMongoConnection() (*mongo.Client, error) {
	// Set up context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB URI - match your docker creds
	fmt.Println("MONGO_URI:", os.Getenv("MONGO_URI"))
	uri := os.Getenv("MONGO_URI")

	// Create client options
	clientOpts := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}
	fmt.Println("Connected to MongoDB!")

	return client, nil
}
