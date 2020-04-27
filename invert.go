package main

import (
	"fmt"
	"net/http"
)

func invert(w http.ResponseWriter, r *http.Request){
	// get multisimensional array from csv file
	rd := csvReader(w, r)

	// calculate the heigth (h) and width (w) of the matrix
	wd := len(rd)
	hg := len(rd[0])

	if hg == 0 || wd == 0 { 
		fmt.Fprint(w, "zero/zero matrix")
		return
	}

	for i := 0; i < hg; i++ {
		for j := 0; j < wd; j++ {

			fmt.Fprint(w, rd[j][i])

			if j == wd -1 {
				fmt.Fprint(w, "\n")
			} else {
				fmt.Fprint(w, ",")
			}
		}
	}
}