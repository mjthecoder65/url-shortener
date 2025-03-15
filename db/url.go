package db

import (
	"context"
	"time"

	"github.com/mjthecoder65/url-shortener/config"
	"github.com/mjthecoder65/url-shortener/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateShortURLParams struct {
	URL string `json:"url" bson:"url"`
}

func (q *Queries) CreateShortURL(ctx context.Context, config *config.Config, arg CreateShortURLParams) (URL, error) {
	collection := q.client.Database("main").Collection("urls")

	var shortCode string
	var err error

	for {

		shortCode, err = utils.GenerateShortCode(config)

		if err != nil {
			return URL{}, nil
		}

		count, err := collection.CountDocuments(ctx, bson.M{"shortCode": shortCode})

		if err != nil {
			return URL{}, nil
		}

		if count == 0 {
			break
		}
	}

	currentTime := time.Now()

	urlDoc := URL{
		URL:         arg.URL,
		ShortCode:   shortCode,
		AccessCount: 0,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	result, err := collection.InsertOne(ctx, urlDoc)

	if err != nil {
		return URL{}, nil
	}

	urlDoc.ID = result.InsertedID.(primitive.ObjectID)
	return urlDoc, nil
}

func (q *Queries) GetShortURL(ctx context.Context, shortCode string) (URL, error) {
	collection := q.client.Database("main").Collection("urls")

	var url URL
	err := collection.FindOne(ctx, bson.M{"shortCode": shortCode}).Decode(&url)

	if err != nil {
		return URL{}, err
	}
	_, err = collection.UpdateOne(ctx, bson.M{"shortCode": shortCode}, bson.M{"$inc": bson.M{"accessCount": 1}})

	if err != nil {
		return URL{}, nil
	}

	return url, nil
}

type UpdateShortURLParams struct {
	ShortCode string `json:"shortCode" bson:"shortCode"`
	URL       string `json:"url" bson:"url"`
}

func (q *Queries) UpdateShortURL(ctx context.Context, arg UpdateShortURLParams) (URL, error) {
	collection := q.client.Database("main").Collection("urls")

	update := bson.M{
		"$set": bson.M{
			"url":       arg.URL,
			"updatedAt": time.Now(),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedShortURL URL

	err := collection.FindOneAndUpdate(
		ctx,
		bson.M{"shortCode": arg.ShortCode},
		update,
		opts,
	).Decode(&updatedShortURL)

	if err != nil {
		return URL{}, err
	}

	return updatedShortURL, nil
}

func (q *Queries) DeleteShortURL(ctx context.Context, shortCode string) error {
	collection := q.client.Database("main").Collection("urls")

	result, err := collection.DeleteOne(ctx, bson.M{"shortCode": shortCode})

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (q *Queries) GetShortURLStats(ctx context.Context, shortCode string) (URL, error) {
	return q.GetShortURL(ctx, shortCode)
}
