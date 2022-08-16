package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func readBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	book, err := getBook(title)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(book)
}
func createBook(w http.ResponseWriter, r *http.Request) {
}

func updateBook(w http.ResponseWriter, r *http.Request) {
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}", createBook).Methods("POST")
	r.HandleFunc("/books/{title}", readBook).Methods("GET")
	r.HandleFunc("/books/{title}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":80", r)
}
