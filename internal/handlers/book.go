package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetBook(c echo.Context) error {
	// slug := c.Param("slug")
	// a, err := h.articleStore.GetBySlug(slug)

	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	// }

	// if a == nil {
	// 	return c.JSON(http.StatusNotFound, utils.NotFound())
	// }

	return c.JSON(http.StatusOK, "Test")
}
