package publicapis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"books.api/internal/common"
	"github.com/labstack/echo"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (t *Handler) GetApi(c echo.Context) error {
	// Call api
	res, err := http.Get(os.Getenv("PUBLICAPIS_URL"))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Get boby response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	// Json to struct
	var root rootResponse
	json.Unmarshal(body, &root)
	os.Exit(0)

	return c.JSON(http.StatusOK, common.Response(root))
}
