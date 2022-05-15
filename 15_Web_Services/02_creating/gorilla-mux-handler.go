package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = "8080"
)

// GetRequest Handler
var GetRequest = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Request Handler")
	},
)

// PostRequest Handler
var PostRequest = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Post Request Handler")
	},
)

// PathVariable Handler
var PathVariable = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprint(w, "Hello, ", name)
	},
)

func main() {

	router := mux.NewRouter()
	router.Handle("/", GetRequest).Methods("GET")
	router.Handle("/post", PostRequest).Methods("POST")
	router.Handle("/about/{name}", PathVariable).Methods("GET", "PUT")
	http.ListenAndServe(Host+":"+Port, router)

}
