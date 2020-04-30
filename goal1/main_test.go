package main

import (
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
