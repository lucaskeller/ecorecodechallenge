package main

import (
	"fmt"
	"net/http"
	"strings"
)

func performEcho(records [][]string) string {
		response := ""

		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}

		return response
}

func echo(w http.ResponseWriter, r *http.Request) {
	// get multisimensional array from csv file
	records := csvReader(w, r)

	rd := performEcho(records)

	fmt.Fprint(w, rd)
}