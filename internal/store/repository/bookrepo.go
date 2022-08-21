package bookrepo

import (
	"context"
	"os"

	"books.api/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStore struct {
	*mongo.Collection
}

func NewMongoDBStore() *MongoDBStore {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_CONNECT")))
	if err != nil {
		panic("failed to connect database")
	}
	collection := client.Database(os.Getenv("MONGO_DB_NAME")).Collection("books")

	return &MongoDBStore{Collection: collection}
}

func (s *MongoDBStore) GetById(bookId string) (entity.Book, error) {
	id, _ := primitive.ObjectIDFromHex(bookId)

	var (
		ctx    = context.Background()
		filter = bson.M{"_id": id}
		result entity.Book
	)

	// Find
	err := s.Collection.FindOne(ctx, filter).Decode(&result)

	return result, err
}

func (s *MongoDBStore) GetAll() ([]entity.Book, error) {

	var (
		ctx    = context.Background()
		filter = bson.M{}
		result []entity.Book
	)

	// Find All
	cursor, err := s.Collection.Find(ctx, filter)
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var item entity.Book
		cursor.Decode(&item)
		result = append(result, item)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, err
}

func (s *MongoDBStore) Add(book *entity.Book) error {
	var ctx = context.Background()

	// Insert
	_, err := s.Collection.InsertOne(ctx, book)
	return err
}

func (s *MongoDBStore) Update(book *entity.Book) error {
	update := bson.M{
		"$set": book,
	}
	var ctx = context.Background()

	// Update
	_, err := s.Collection.UpdateByID(ctx, book.Id, update)
	return err
}

func (s *MongoDBStore) Delete(bookId string) error {
	id, _ := primitive.ObjectIDFromHex(bookId)

	var (
		ctx    = context.Background()
		filter = bson.M{"_id": id}
	)

	// Delete
	_, err := s.Collection.DeleteOne(ctx, filter)
	return err
}
