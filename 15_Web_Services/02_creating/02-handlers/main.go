package main

import (
	"fmt"
	"log"
	"net/http"
)

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "default endpoint")
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World")
}

func main() {
	// Routing
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/hello", helloHandler)

	log.Println("Starting server at 8000 ...")

	// Starting server - default is localhost
	// log.Fatal(http.ListenAndServe(":8000", nil))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

/*
USage:
	~curl http://localhost:8000/
	default endpoint
	~curl http://localhost:8000/hello
	Hello World
	~curl http://localhost:8000/hello2
	default endpoint

*/
