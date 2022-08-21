package main

import (
	"log"

	"books.api/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("local.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	router.New()
}
