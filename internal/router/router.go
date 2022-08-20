package router

import (
	"context"

	"books.api/internal/domain"
	"books.api/internal/store"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//create groups
	// adminGroup := e.Group("/admin")

	//set all middlewares
	// middlewares.SetMainMiddleWares(e)
	// middlewares.SetAdminMiddlewares(adminGroup)

	//set main routes
	// api.MainGroup(e)

	//set groupRoutes
	// api.AdminGroup(adminGroup)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		panic("failed to connect database")
	}
	collection := client.Database("bookstore").Collection("books")
	// gormStore := store.NewGormStore(db)
	mongoStore := store.NewMongoDBStore(collection)

	handler := domain.NewHandler(mongoStore)

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.POST("/books", handler.AddBook)
		// v1.GET("/members", api.GetMembers())
		// v1.GET("/members/:id", api.GetMember())
	}
	return e
}
