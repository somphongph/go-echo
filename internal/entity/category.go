package entity

import (
	"time"

	"books.api/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      model.Locale       `json:"name"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
