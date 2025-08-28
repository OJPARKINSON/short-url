package main

import (
	"log"
	"net/http"

	"github.com/ojparkinson/shortUrl/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mux := http.NewServeMux()

	mux.Handle("/shorten", &handlers.ShortenHandler{})

	log.Fatal(http.ListenAndServe(":8090", mux))
}
