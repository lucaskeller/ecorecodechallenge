package main

import (
	"fmt"
	"net/http"
	"strings"
)

func performFlatten(records [][]string) string {
		var response string
		for i, row := range records {
			response = response + strings.Join(row, ",")
			if len(records) != i + 1 {
				response = response + ","
			} else {
				response = response + "\n"
			}
		}

		return response
}

func flatten(w http.ResponseWriter, r *http.Request){
		// get multisimensional array from csv file
		records := csvReader(w, r)

		response := performFlatten(records)
		
		fmt.Fprint(w, response)
}