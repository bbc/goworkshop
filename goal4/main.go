package main

import (
	"log"
	"net/http"
)

// EncodeRequest represents a request to encode
type EncodeRequest struct {
	URL string
}

// EncodeResponse represents the result of encoding
type EncodeResponse struct {
	ShortURL string
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

}
