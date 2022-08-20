package router

import (
	"books.api/internal/handler/book"
	"books.api/internal/store"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Routes
	api := e.Group("/api")

	// books
	bookHandler := book.NewHandler(store.NewMongoDBStore())
	books := api.Group("/v1/books")
	{
		books.POST("", bookHandler.AddBook)
		// v1.GET("/members", api.GetMembers())
		// v1.GET("/members/:id", api.GetMember())
	}

	// authors
	// authorHandler := book.NewHandler(store.NewMongoDBStore())
	authors := api.Group("/v1/authors")
	{
		authors.POST("", bookHandler.AddBook)
		// v1.GET("/members", api.GetMembers())
		// v1.GET("/members/:id", api.GetMember())
	}

	// categories
	// categories := api.Group("/categories")
	// {
	// 	categories.POST("", domain.AddBook)
	// 	// v1.GET("/members", api.GetMembers())
	// 	// v1.GET("/members/:id", api.GetMember())
	// }

	e.Logger.Fatal(e.Start(":1323"))

	return e
}
