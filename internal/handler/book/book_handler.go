package book

import (
	"net/http"
	"time"

	"books.api/internal/common"
	"books.api/internal/entity"
	"books.api/internal/model"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type storer interface {
	Add(*entity.Book) error
}

type Handler struct {
	store storer
}

func NewHandler(store storer) *Handler {
	return &Handler{store: store}
}

func (t *Handler) AddBook(c echo.Context) error {
	bookRequest := new(bookRequest)
	if err := c.Bind(&bookRequest); err != nil {
		return err
	}

	book := &entity.Book{
		Id:        primitive.NewObjectID(),
		Isbn:      bookRequest.Isbn,
		Name:      model.Locale{En: bookRequest.NameEn, Th: bookRequest.NameTh},
		Title:     model.Locale{En: bookRequest.TitleEn, Th: bookRequest.TitleTh},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := t.store.Add(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")

		return err
	}

	return c.JSON(http.StatusOK, common.Response(book))
}
