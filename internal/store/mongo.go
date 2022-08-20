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

// func (s *MongoDBStore) Get(book *domain.Book) error {
// 	_, err := s.Collection.Find(book ,"")
// 	return err
// }

// func (s *MongoDBStore) Gets(book *domain.Book) error {
// 	_, err := s.Collection.InsertOne(context.Background(), book)
// 	return err
// }

func (s *MongoDBStore) Add(book *domain.Book) error {
	_, err := s.Collection.InsertOne(context.Background(), book)
	return err
}

// func (s *MongoDBStore) Update(book *domain.Book) error {
// 	_, err := s.Collection.UpdateOne(context.Background(), book)
// 	return err
// }

func (s *MongoDBStore) Delete(book *domain.Book) error {
	_, err := s.Collection.DeleteOne(context.Background(), book)
	return err
}
