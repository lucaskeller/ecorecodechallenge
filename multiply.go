package main

import (
	"strconv"
	"fmt"
	"net/http"
)

func multiply(w http.ResponseWriter, r *http.Request){
	// get multisimensional array from csv file
	rd := csvReader(w, r)

	// calculate the heigth (h) and width (w) of the matrix
	wd := len(rd)
	hg := len(rd[0])

	if hg == 0 || wd == 0 { 
		fmt.Fprint(w, 0)
		return
	}

	p := 0

	for i := 0; i < hg; i++ {
		for j := 0; j < wd; j++ {

			v := rd[i][j]

			ival, err := strconv.Atoi(v)

			if err == nil {
				if p == 0 {
					p = ival
				} else {
					p = p * ival
				}
			} else {
				fmt.Fprint(w, "Error converting string into int")
			}
		}
	}

	fmt.Fprint(w, p)
}