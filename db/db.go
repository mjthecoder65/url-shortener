package db

import "go.mongodb.org/mongo-driver/mongo"

type Queries struct {
	client *mongo.Client
}

func New(client *mongo.Client) *Queries {
	return &Queries{
		client: client,
	}
}
