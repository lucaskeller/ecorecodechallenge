package main

import (
	"strconv"
	"fmt"
	"net/http"
)

func performSum(records [][]string) string {
	s := 0

	// calculate the heigth (h) and width (w) of the matrix
	wd := len(records)
	hg := len(records[0])

	if hg == 0 || wd == 0 { 
		return string(s)
	}

	for i := 0; i < hg; i++ {
		for j := 0; j < wd; j++ {

			v := records[i][j]

			ival, err := strconv.Atoi(v)

			if err == nil {
				s = s + ival
			} else {
				return "Error converting string into int"
			}
		}
	}

	return string(s)
}

func sum(w http.ResponseWriter, r *http.Request){
	// get multisimensional array from csv file
	records := csvReader(w, r)

	s := performSum(records)

	fmt.Fprint(w, s)
}