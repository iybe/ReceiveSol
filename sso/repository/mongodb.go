package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientMongoDB struct {
	Client          *mongo.Client
	DatabaseName    string
	CollectionUsers string
}

type User struct {
	ID       string `bson:"_id,omitempty"`
	User     string `bson:"user"`
	Password string `bson:"password"`
}

func CreateClient(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}

func (c *ClientMongoDB) FindUser(user string) (*User, error) {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionUsers)

	var newUser User
	err := collection.FindOne(context.Background(), bson.M{"user": user}).Decode(&newUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Printf("Error finding user by username: %v\n", err)
		return nil, err
	}

	return &newUser, nil
}

func (client *ClientMongoDB) AddUser(user User) (*mongo.InsertOneResult, error) {
	collection := client.Client.Database(client.DatabaseName).Collection(client.CollectionUsers)
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
