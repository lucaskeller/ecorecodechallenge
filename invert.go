package main

import (
	"fmt"
	"net/http"
)

func performInvert(records [][]string) string {
	response := ""

	// calculate the heigth (h) and width (w) of the matrix
	wd := len(records)
	hg := len(records[0])

	if hg == 0 || wd == 0 { 
		return "zero/zero matrix"
	}

	for i := 0; i < hg; i++ {
		for j := 0; j < wd; j++ {

			response = response + string(records[j][i])

			if j == wd -1 {
				response = response + "\n"
			} else {
				response = response + ","
			}
		}
	}

	return response
}

func invert(w http.ResponseWriter, r *http.Request){
	// get multisimensional array from csv file
	records := csvReader(w, r)

	rd := performInvert(records)

	fmt.Fprint(w, rd)
}