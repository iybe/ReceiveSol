package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	ID        string `bson:"_id, omitempty"`
	PublicKey string `bson:"publicKey"`
	UserId    string `bson:"userId"`
	Nickname  string `bson:"nickname"`
}

func (c *ClientMongoDB) AddAccount(account Account) (*Account, error) {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionAccount)
	result, err := collection.InsertOne(context.Background(), account)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	newAccount := &Account{
		ID:        id.Hex(),
		PublicKey: account.PublicKey,
		UserId:    account.UserId,
		Nickname:  account.Nickname,
	}

	return newAccount, nil
}

func (c *ClientMongoDB) GetAccountByPublicKey(publicKey string) (*Account, error) {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionAccount)
	var newAccount Account
	err := collection.FindOne(context.Background(), bson.M{"publicKey": publicKey}).Decode(&newAccount)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &newAccount, nil
}

func (c *ClientMongoDB) ListAccountByUserId(userId string) ([]Account, error) {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionAccount)
	var accounts []Account
	cur, err := collection.Find(context.Background(), bson.M{"userId": userId})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var account Account
		err := cur.Decode(&account)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
