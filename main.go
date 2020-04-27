package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/Users/lucaskeller/projects/ecore/codechallenge/matrix.csv' "localhost:8181/echo"

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// sytem routes
	router.HandleFunc("/sum", sum)
	router.HandleFunc("/echo", echo)
	router.HandleFunc("/invert", invert)
	router.HandleFunc("/flatten", flatten)
	router.HandleFunc("/multiply", multiply)

	http.ListenAndServe(":8080", router)
}
