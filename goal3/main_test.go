package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestHandlerOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	hnd := http.HandlerFunc(handler)

	hnd.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("incorrect status code: got %v want %v", status, http.StatusOK)
	}
}

func TestHandlerAcceptsURL(t *testing.T) {
	var jsonReq = []byte(`{"URL":"https://www.bbc.co.uk/iplayer"}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	hnd := http.HandlerFunc(handler)

	hnd.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("incorrect status code: got %v want %v", status, http.StatusBadRequest)
	}

	jsonBody := json.NewDecoder(rec.Body)
	var encRes EncodeResponse
	jsonBody.Decode(&encRes)

	u := regexp.MustCompile(`localhost:8080/\d`)
	if !u.MatchString(encRes.ShortURL) {
		t.Errorf("URL doesn't match expected format: got %s want format as %s", encRes.ShortURL, u.String())
	}
}

func TestHandlerRejectsMissingURL(t *testing.T) {
	var jsonReq = []byte(`{"URL":""}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	hnd := http.HandlerFunc(handler)

	hnd.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("incorrect status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestHandlerRejectsMissingBody(t *testing.T) {
	var jsonReq = []byte(``)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	hnd := http.HandlerFunc(handler)

	hnd.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("incorrect status code: got %v want %v", status, http.StatusBadRequest)
	}
}
