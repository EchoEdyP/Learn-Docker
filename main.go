package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})
	http.ListenAndServe("localhost:8080", nil)
}
