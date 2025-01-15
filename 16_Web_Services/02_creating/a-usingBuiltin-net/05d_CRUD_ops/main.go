package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var data = map[string]string{
	"hello": "world",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request")
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, data[r.URL.Path])
			log.Println("GET request handled")
		case "POST":
			key := r.URL.Path
			value, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			data[key] = string(value)
			fmt.Fprintf(w, "key: %s, value: %s\n", key, value)
			log.Println("POST request handled")
		case "PUT":
			key := r.URL.Path
			value, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			data[key] = string(value)
			fmt.Fprintf(w, "key: %s, value: %s\n", key, value)
			log.Println("PUT request handled")
		case "DELETE":
			delete(data, r.URL.Path)
			fmt.Fprintf(w, "key: %s deleted\n", r.URL.Path)
			log.Println("DELETE request handled")
		default:
			fmt.Fprintf(w, "Method not allowed\n")
			log.Println("Method not allowed")
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
