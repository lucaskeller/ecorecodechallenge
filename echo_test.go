package main

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"bytes"
	"net/http/httptest"
	"testing"
)


func TestEcho(t *testing.T) {
	//
	// CREATING THE FILE TO TEST

	// creating the body content to send a file via multipar/form-data
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "matrix.csv")
	
	// read some local file
	bs, err := ioutil.ReadFile("matrix.csv")
	if err != nil {
		t.Errorf("Error reading local file for testing: %v", err)
	}

	// write the file to the pipe
	part.Write(bs)

	// finish multipart adding file boundaries
	writer.Close()

	// create a new http.Request fopr testing purpose
	// using post to have how to set the body content
	req := httptest.NewRequest("POST", "/echo", body)

	// set the content type, with previous boundaries
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// returns a *ResponseRecorder
	rr := httptest.NewRecorder()

	echo(rr, req)

	resp := rr.Result()
	b, _ := ioutil.ReadAll(resp.Body)

	// get echo response
	er := string(b)

	// start implementing test cases

	// get the file contents
	csvf := csvReader(rr, req)

	fmt.Println(csvf)
	fmt.Println(er)

	//
	// 1. check string length
	// if er
}