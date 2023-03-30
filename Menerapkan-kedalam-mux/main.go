package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Use(Logging)

	r.HandleFunc("/foo", foo).Methods("GET")
	r.HandleFunc("/bar", bar).Methods("GET")

	http.ListenAndServe(":8080", r)
}