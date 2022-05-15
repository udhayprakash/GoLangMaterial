package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println(id)

	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/product/{id:[0-9]+}", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

/*
OUTPUT:
	~curl http://localhost:8000/
	404 page not found

	~curl http://localhost:8000/product
	404 page not found

	~curl http://localhost:8000/product/1
	Gorilla!

	~curl http://localhost:8000/product/2
	Gorilla!


*/
