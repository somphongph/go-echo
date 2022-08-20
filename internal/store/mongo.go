package store

import (
	"context"

	"books.api/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBStore struct {
	*mongo.Collection
}

func NewMongoDBStore(col *mongo.Collection) *MongoDBStore {
	return &MongoDBStore{Collection: col}
}

func (s *MongoDBStore) Add(book *domain.Book) error {
	_, err := s.Collection.InsertOne(context.Background(), book)
	return err
}
