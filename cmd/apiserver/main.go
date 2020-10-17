package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/add", addShortURLs) // POST
	http.HandleFunc("/", redirect) // GET

	log.Fatal(http.ListenAndServe(":8080", nil))
}
