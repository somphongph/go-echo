package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name"`
	Title     string             `json:"title"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
