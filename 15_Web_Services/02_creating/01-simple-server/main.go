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

	// // Method 1 - windows security alert comes here
	// log.Fatal(http.ListenAndServe(":8000", nil))

	// // Method 2 - explicitly mentioning host. Then no prompt
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))
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
