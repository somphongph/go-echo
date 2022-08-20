package book

import "books.api/internal/model"

type bookRequest struct {
	Name  model.Locale `json:"name"`
	Title model.Locale `json:"title"`
}
