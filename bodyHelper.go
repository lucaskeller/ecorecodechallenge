package main

import (
	"mime/multipart"
	"bytes"
)

// bs: byte slice
// fn: filename
func bodyHelper(bs []byte, fn string) (*bytes.Buffer, *multipart.Writer) {
	// creating the body content to send a file via multipar/form-data

	// allocate a bytes.Buffer
	body := new(bytes.Buffer)
	// create the multipart writer
	writer := multipart.NewWriter(body)
	// create the form-data header for the file
	part, _ := writer.CreateFormFile("file", fn)

	// write the file to the pipe
	part.Write(bs)

	// finish multipart adding file boundaries
	writer.Close()

	return body, writer
}