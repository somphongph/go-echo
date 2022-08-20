package main

import (
	"books.api/internal/router"
)

func main() {
	// create a new echo instance
	e := router.New()

	e.Logger.Fatal(e.Start(":1323"))
}
