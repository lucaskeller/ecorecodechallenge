package main

import (
	"strconv"
	"fmt"
	"net/http"
)

//
// this program have a problem with big values
// I tried to figure out how to deal with big.Int without success
// I'll continue trying, but for now that's what I have

func performMultiply(records [][]string) string {
// calculate the heigth (h) and width (w) of the matrix
	wd := len(records)
	hg := len(records[0])

	if hg == 0 || wd == 0 { 
		return "0"
	}

	p := 0

	for i := 0; i < hg; i++ {
		for j := 0; j < wd; j++ {
			v := records[i][j]

			ival, err := strconv.Atoi(v)

			if err == nil {
				if p == 0 {
					p = ival
				} else {
					p = p * ival
				}
			} else {
				return "Problems converting string to int"
			}
		}
	}

	return string(p)
}

func multiply(w http.ResponseWriter, r *http.Request){
	// get multisimensional array from csv file
	records := csvReader(w, r)

	p := performMultiply(records)

	fmt.Fprint(w, p)
}