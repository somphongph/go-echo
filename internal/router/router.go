package router

import (
	"os"

	"books.api/internal/handler/book"
	"books.api/internal/handler/publicapis"
	"books.api/internal/store/cache"
	"books.api/internal/store/repository"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Routes
	api := e.Group("/api")

	// books
	bookHandler := book.NewHandler(repository.NewMongoDBStore(), cache.InitCache())
	books := api.Group("/v1/books")
	{
		books.GET("/:id", bookHandler.GetById)
		books.GET("", bookHandler.GetAll)
		books.POST("", bookHandler.AddBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}

	publicapisHandler := publicapis.NewHandler()
	publicapis := api.Group("/v1/publicapis")
	{
		publicapis.GET("", publicapisHandler.GetApi)
	}

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

	return e
}
