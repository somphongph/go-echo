package store

import (
	"context"

	"books.api/internal/handlers/book"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStore struct {
	*mongo.Collection
}

func NewMongoDBStore() *MongoDBStore {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		panic("failed to connect database")
	}
	collection := client.Database("bookstore").Collection("books")

	return &MongoDBStore{Collection: collection}
}

// func (s *MongoDBStore) Get(book *domain.Book) error {
// 	_, err := s.Collection.Find(book ,"")
// 	return err
// }

// func (s *MongoDBStore) Gets(book *domain.Book) error {
// 	_, err := s.Collection.InsertOne(context.Background(), book)
// 	return err
// }

func (s *MongoDBStore) Add(book *book.Book) error {
	_, err := s.Collection.InsertOne(context.Background(), book)
	return err
}

// func (s *MongoDBStore) Update(book *domain.Book) error {
// 	_, err := s.Collection.UpdateOne(context.Background(), book)
// 	return err
// }

func (s *MongoDBStore) Delete(book *book.Book) error {
	_, err := s.Collection.DeleteOne(context.Background(), book)
	return err
}
