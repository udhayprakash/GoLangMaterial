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
	// http.DefaultServeMux.HandleFunc("/page2", Page2)
	// http.DefaultServeMux.HandleFunc("/", Index)

	// Method 3
	mux := http.NewServeMux()
	mux.HandleFunc("/page2", Page2)
	mux.HandleFunc("/", Index)

	// Starting the server
	log.Println("Starting server at 3000 ...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func Page2(res http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf(res, "Page2 endpoint")
	res.Write([]byte("<h1>Welcome to Page2!</h1>"))
}

func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Index endpoint")
}
