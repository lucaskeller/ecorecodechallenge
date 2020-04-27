package main

import (
	"encoding/csv"
	"io/ioutil"
	"bytes"
	"net/http/httptest"
	"testing"
)


func TestInvert(t *testing.T) {
	//
	// CREATING THE FILE TO TEST

	// 1. read some local file to run throught service an locally
	bs, err := ioutil.ReadFile("matrix.csv")
	if err != nil {
		t.Errorf("Error reading local file for testing: %v", err)
	}

	//
	// SEND THE FILE TO SERVER

	// create the body
	body, writer := bodyHelper(bs, "matrix.csv")

	// create a new http.Request fopr testing purpose
	// using post to have how to set the body content
	req := httptest.NewRequest("POST", "/invert", body)

	// set the content type, with previous boundaries
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// returns a *ResponseRecorder
	rr := httptest.NewRecorder()

	invert(rr, req)

	// get the endpoint response
	resp := rr.Result()
	b, _ := ioutil.ReadAll(resp.Body)

	// byte to string
	sr := string(b)


	//
	// NOW RUN THE FUNCTION DIRECTLY TO TEST
	// convert byte slice to io.Reader
	bsior := bytes.NewReader(bs)
	records, err := csv.NewReader(bsior).ReadAll()
	if err != nil {
		t.Errorf("Error reading csv file")
	}
	tr := performInvert(records)


	if sr != tr {
		t.Errorf("Expected the same result as %v, got %v", sr, tr)
	}
}