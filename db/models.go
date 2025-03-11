package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type URL struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	URL         string             `json:"url" bson:"url"`
	ShortCode   string             `json:"shortCode" bson:"shortCode"`
	AccessCount int64              `json:"accessCount" bson:"accessCount"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type ShortURLStats struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	URL         string             `json:"url" bson:"url"`
	ShortCode   string             `json:"shortCode" bson:"shortCode"`
	AccessCount int64              `json:"accessCount" bson:"accessCount"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}
