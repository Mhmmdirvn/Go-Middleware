package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", Logging(foo))
	http.HandleFunc("/bar", Logging(bar))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}