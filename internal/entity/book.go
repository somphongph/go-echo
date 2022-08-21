package entity

import (
	"time"

	"books.api/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Isbn      string             `json:"isbn"`
	Name      model.Locale       `json:"name"`
	Title     model.Locale       `json:"title"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
