package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", IndexHandlerFunc)

	// Starting the server
	log.Println("Starting server at 8000 ...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func IndexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// fmt.Fprint(w, "Welcome to my index Page!")
	w.Write([]byte("<h1>Welcome to my index Page!</h1>"))
}
