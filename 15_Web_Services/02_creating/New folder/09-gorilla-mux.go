package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "up and running")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPosts).Methods("POST")

	log.Println("Server Listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}

func getPosts() {
	fmt.Println("getPosts - start")
}

func addPosts() {
	fmt.Println("addPosts - start")
}
