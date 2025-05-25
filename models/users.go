package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	collectionName = "users"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (client *Client) CreateUser(user *User) error {
	collection := client.client.Database(client.dbName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), user)
	return err
}

func (client *Client) GetUsers() ([]User, error) {
	collection := client.client.Database(client.dbName).Collection(collectionName)
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var results []User
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (client *Client) GetUserByID(id string) (*User, error) {
	collection := client.client.Database(client.dbName).Collection(collectionName)
	var result User
	err := collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&result)
	return &result, err
}
