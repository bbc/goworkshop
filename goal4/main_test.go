package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestHandlerNotImplemented(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	hnd := http.HandlerFunc(handler)

	hnd.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusNotImplemented {
		t.Errorf("incorrect status code: got %v want %v", status, http.StatusNotImplemented)
	}
}

func TestHandleDuplicateURL(t *testing.T) {
	var jsonReq = []byte(`{"URL":"https://www.bbc.co.uk/iplayer"}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	hnd := http.HandlerFunc(handler)

	hnd.ServeHTTP(rec, req)

	jsonBody := json.NewDecoder(rec.Body)
	var encRes1 EncodeResponse
	jsonBody.Decode(&encRes1)

	req2, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec2 := httptest.NewRecorder()
	hnd2 := http.HandlerFunc(handler)

	hnd2.ServeHTTP(rec2, req2)

	jsonBody2 := json.NewDecoder(rec2.Body)
	var encRes2 EncodeResponse
	jsonBody2.Decode(&encRes2)

	if encRes1.ShortURL != encRes2.ShortURL {
		t.Errorf("duplicate URL incorrectly creates new ShortURL: %s and %s", encRes1.ShortURL, encRes2.ShortURL)
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
