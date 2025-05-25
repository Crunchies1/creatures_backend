package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	client *mongo.Client
	dbName string
}

func NewClient(client *mongo.Client, dbName string) *Client {
	return &Client{client: client, dbName: dbName}
}
