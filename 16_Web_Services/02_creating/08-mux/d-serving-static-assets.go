package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	staticHandler := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", staticHandler))

	// Starting the server
	log.Println("Starting server at 3000 ...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func redirect(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, 301)
	}
}

func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Index endpoint")
}
