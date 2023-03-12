package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientMongoDB struct {
	Client            *mongo.Client
	DatabaseName      string
	CollectionUsers   string
	CollectionAccount string
	CollectionLink    string
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
