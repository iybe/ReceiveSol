package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID                  string `bson:"_id,omitempty"`
	Username            string `bson:"username"`
	PermaLinkUrl        string `bson:"permaLinkUrl,omitempty"`
	PermaLinkRecipient  string `bson:"permaLinkRecipient,omitempty"`
	PermaLinkExpiration int64  `bson:"permaLinkExpiration,omitempty"`
}

func (c *ClientMongoDB) AddUser(user User) (*User, error) {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionUsers)
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	newUser := &User{
		ID:       id.Hex(),
		Username: user.Username,
	}

	return newUser, nil
}

func (c *ClientMongoDB) GetUser(username string) (*User, error) {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionUsers)
	var user User
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (c *ClientMongoDB) GetUserById(id string) (*User, error) {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionUsers)

	obj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user User
	err = collection.FindOne(context.Background(), bson.M{"_id": obj}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *ClientMongoDB) UpdateUserPermaLink(id string, permaLinkUrl string, permaLinkRecipient string, permaLinkExpiration int64) error {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionUsers)
	obj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	var user User
	err = collection.FindOne(context.Background(), bson.M{"_id": obj}).Decode(&user)
	if err != nil {
		return err
	}

	user.PermaLinkExpiration = permaLinkExpiration
	user.PermaLinkRecipient = permaLinkRecipient
	user.PermaLinkUrl = permaLinkUrl

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientMongoDB) SearchPermaLinkUrl(permaLinkUrl string) (*User, error) {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionUsers)
	var user User
	err := collection.FindOne(context.Background(), bson.M{"permaLinkUrl": permaLinkUrl}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
