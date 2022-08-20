package domain

import (
	"net/http"

	"github.com/labstack/echo"
)

func (Book) TableName() string {
	return "books"
}

type storer interface {
	Add(*Book) error
}

type Handler struct {
	store storer
}

func NewHandler(store storer) *Handler {
	return &Handler{store: store}
}

func (t *Handler) AddBook(c echo.Context) error {
	var book Book

	err := t.store.Add(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Test")

		return err
	}

	return c.JSON(http.StatusOK, "Test")
	// c.JSON(http.StatusCreated, map[string]interface{}{
	// 	"ID": todo.ID,
	// })
}
