package book

import (
	"encoding/json"
	"net/http"
	"time"

	// "books.api/internal/cache"
	"books.api/internal/common"
	"books.api/internal/entity"
	"books.api/internal/model"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type storer interface {
	GetById(string) (entity.Book, error)
	GetAll() ([]entity.Book, error)
	Add(*entity.Book) error
	Update(*entity.Book) error
	Delete(string) error
}

type cached interface {
	Cache(string, interface{}, time.Duration) error
	ShortCache(string, interface{}) error
	LongCache(string, interface{}) error
}

type Handler struct {
	store storer
	cache cached
}

func NewHandler(store storer, cache cached) *Handler {
	return &Handler{store: store, cache: cache}
}

// Get item book
func (t *Handler) GetById(c echo.Context) error {
	id := c.Param("id")
	var book entity.Book

	book, err := t.store.GetById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, common.ResponseNotFound())
	}

	// Cached
	data, _ := json.Marshal(book)
	t.cache.ShortCache("xxxx", data)

	return c.JSON(http.StatusOK, common.Response(book))
}

// Get all book
func (t *Handler) GetAll(c echo.Context) error {
	var books []entity.Book

	books, err := t.store.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.ResponseFailed())
	}

	return c.JSON(http.StatusOK, common.Response(books))
}

// Add book
func (t *Handler) AddBook(c echo.Context) error {
	bookRequest := new(bookRequest)
	if err := c.Bind(&bookRequest); err != nil {
		return err
	}

	// Bind object
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
		c.JSON(http.StatusInternalServerError, common.ResponseFailed())

		return err
	}

	return c.JSON(http.StatusOK, common.Response(book))
}

// Update book
func (t *Handler) UpdateBook(c echo.Context) error {
	params := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(params)

	bookRequest := new(bookRequest)
	if err := c.Bind(&bookRequest); err != nil {
		return err
	}

	// Bind object
	book := &entity.Book{
		Id:        id,
		Isbn:      bookRequest.Isbn,
		Name:      model.Locale{En: bookRequest.NameEn, Th: bookRequest.NameTh},
		Title:     model.Locale{En: bookRequest.TitleEn, Th: bookRequest.TitleTh},
		UpdatedAt: time.Now(),
	}

	err := t.store.Update(book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.ResponseFailed())
	}

	return c.JSON(http.StatusOK, common.Response(book))
}

// Delete book
func (t *Handler) DeleteBook(c echo.Context) error {
	bookId := c.Param("id")

	bookRequest := new(bookRequest)
	if err := c.Bind(&bookRequest); err != nil {
		return err
	}

	err := t.store.Delete(bookId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.ResponseFailed())
	}

	return c.JSON(http.StatusOK, common.Response(nil))
}
