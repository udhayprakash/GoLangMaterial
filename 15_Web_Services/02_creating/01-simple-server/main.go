package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "default endpoint")
	})

	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello World")
	})

	log.Println("Starting server at 8000 ...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// http://localhost:8000/
// http://localhost:8000/hello
