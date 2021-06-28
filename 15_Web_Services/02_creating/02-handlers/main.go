package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Routing
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/hello", helloHandler)

	log.Println("Starting server at 8000 ...")

	// Starting server - default is localhost
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "default endpoint")
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World")
}

// http://localhost:8000/
// http://localhost:8000/hello
