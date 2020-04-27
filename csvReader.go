package main

import (
	"fmt"
	"net/http"
	"encoding/csv"
)

func csvReader(w http.ResponseWriter, r *http.Request) [][]string {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return nil
		}
		defer file.Close()

		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return nil
		}

		return records
}