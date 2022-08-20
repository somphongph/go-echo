package router

import (
	"books.api/internal/handlers/book"
	"books.api/internal/store"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	domain := book.NewHandler(store.NewMongoDBStore())

	// Routes
	api := e.Group("/api")

	// books
	books := api.Group("/v1/books")
	{
		books.POST("", domain.AddBook)
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
