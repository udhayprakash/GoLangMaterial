package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Method 1 -- inherently using default mux
	// http.HandleFunc("/page2", Page2)
	// http.HandleFunc("/", Index)

	// Method 2 -- explicitly mentioning
	http.DefaultServeMux.HandleFunc("/page2", Page2)
	http.DefaultServeMux.HandleFunc("/", Index)

	// Starting the server
	log.Println("Starting server at 3000 ...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func Page2(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Page2 endpoint")
}

func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Index endpoint")
}
