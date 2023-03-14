package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	ID             string     `bson:"_id,omitempty"`
	Nickname       string     `bson:"nickname,omitempty"`
	UserId         string     `bson:"userId,omitempty"`
	AccountId      string     `bson:"accountId,omitempty"`
	Link           string     `bson:"link,omitempty"`
	Reference      string     `bson:"reference,omitempty"`
	Recipient      string     `bson:"recipient,omitempty"`
	Network        string     `bson:"network,omitempty"`
	ExpectedAmount float64    `bson:"expectedAmount,omitempty"`
	AmountReceived float64    `bson:"amountReceived,omitempty"`
	Status         string     `bson:"status,omitempty"`
	CreatedAt      *time.Time `bson:"createdAt,omitempty"`
	ReceivedAt     string     `bson:"receivedAt,omitempty"`
	Expiration     int64      `bson:"expiration,omitempty"`
	Expired        bool       `bson:"expired,omitempty"`
	IsPermaLink    bool       `bson:"isPermaLink,omitempty"`
	Code           string     `bson:"code,omitempty"`
}

func (c *ClientMongoDB) CreateLink(link Link) (*Link, error) {
	now := time.Now()
	link.CreatedAt = &now

	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionLink)
	result, err := collection.InsertOne(context.Background(), link)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	newLink := &Link{
		ID:             id.Hex(),
		Reference:      link.Reference,
		Recipient:      link.Recipient,
		Network:        link.Network,
		Nickname:       link.Nickname,
		ExpectedAmount: link.ExpectedAmount,
		Status:         link.Status,
		CreatedAt:      link.CreatedAt,
	}

	return newLink, nil
}

func (c *ClientMongoDB) UpdateLinkStatus(id string, status string) error {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionLink)

	obj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": obj}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientMongoDB) SearchByReference(reference string) (*Link, error) {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionLink)

	var link Link
	err := collection.FindOne(context.Background(), bson.M{"reference": reference}).Decode(&link)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (c *ClientMongoDB) ListLink(userId string, status string, network string, recipient string, permalink bool) ([]Link, error) {
	collection := c.Client.Database(c.DatabaseName).Collection(c.CollectionLink)

	var links []Link
	filter := bson.M{"userId": userId}
	if permalink {
		filter["isPermaLink"] = true
	} else {
		filter["isPermaLink"] = bson.M{"$ne": true}
	}
	if status != "" {
		filter["status"] = status
	}
	if network != "" {
		filter["network"] = network
	}
	if recipient != "" {
		filter["recipient"] = recipient
	}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var link Link
		err := cur.Decode(&link)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return links, nil
}
