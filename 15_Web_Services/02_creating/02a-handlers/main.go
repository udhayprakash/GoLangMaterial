package main

import (
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	Name string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is ", h.Name)
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is Dog!")
}

func main() {
	h := Handler{
		Name: "Cat",
	}
	// http.Handle accepts an interface with the
	// method ServeHTTP(w http.ResponseWriter, r *http.Request)
	http.Handle("/cat", h)

	// http.HandleFunc accepts an http.HandleFunc - a function
	// that looks like Name(w http.ResponseWriter, r *http.Request)
	http.HandleFunc("/dog", HandlerFunc)

	log.Println("Starting server at 8000 ...")

	// Starting the server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
