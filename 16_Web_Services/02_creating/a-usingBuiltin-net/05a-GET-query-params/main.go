package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8090", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// keys, ok := r.URL.Query().Get("key")
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		// fmt.Fprintf(w, "Url Param 'key' is missing")
		http.Error(w, "The key query parameter is missing", http.StatusBadRequest)
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]
	fmt.Fprintf(w, "Url Param 'key' is: "+string(key))
}

// http://localhost:8080/?key=asda
