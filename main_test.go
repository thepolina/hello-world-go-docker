package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloWorld)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if rr.Body.String() != "hello world" {
		t.Errorf("handler returned unexpected body: got %v want 'hello world'",
			rr.Body.String())
	}
}
