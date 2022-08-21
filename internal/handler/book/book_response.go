package book

import (
	"books.api/internal/model"
)

type bookResponse struct {
	Isbn  string       `json:"isbn"`
	Name  model.Locale `json:"name"`
	Title model.Locale `json:"title"`
}
