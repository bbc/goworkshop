package main

import (
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

}

func handler(w http.ResponseWriter, r *http.Request) {

}
