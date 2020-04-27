package main

import (
	"net/http/httptest"
	"net/http"
	"testing"
)


func TestEcho(t *testing.T) {
    req, err := http.NewRequest("GET", "/echo", nil)
    if err != nil {
        t.Fatal(err)
		}
		req.Header.Add("Content-Type", "multipart/form-data")
		req.Body.Add("file", "@/Users/lucaskeller/projects/ecore/codechallenge/matrix.csv")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(echo)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := ""
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}