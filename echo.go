package main

import (
	"fmt"
	"net/http"
	"strings"
)

func echo(w http.ResponseWriter, r *http.Request) {
		records := csvReader(w, r)

		var response string

		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}

		fmt.Fprint(w, response)
}