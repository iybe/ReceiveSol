package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID                 string `bson:"_id,omitempty"`
	Username           string `bson:"username"`
	RecipientPermaLink string `bson:"recipientPermaLink"`
	NetworkPermaLink   string `bson:"networkPermaLink"`
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

func (c *ClientMongoDB) SetPermaLink(userId, recipient, network string) error {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionUsers)
	obj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	collection.FindOneAndUpdate(context.Background(), bson.M{"_id": obj}, bson.M{"$set": bson.M{"recipientPermaLink": recipient, "networkPermaLink": network}})

	return nil
}
