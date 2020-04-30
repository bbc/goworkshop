package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
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

func TestHandlerRespondsWithURL(t *testing.T) {
	jsonReq := []byte(`{"URL":"https://www.bbc.co.uk/"}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	hnd := http.HandlerFunc(handler)

	hnd.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("incorrect status code: got %v want %v", status, http.StatusOK)
	}

	expected := "GOT: https://www.bbc.co.uk/"
	if rec.Body.String() != expected {
		t.Errorf("incorrect response: got %s want %s", rec.Body.String(), expected)
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
