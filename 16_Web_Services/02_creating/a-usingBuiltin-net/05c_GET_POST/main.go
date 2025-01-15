package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/hello/", hello)
	http.ListenAndServe("localhost:8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// Get the name and age from the request body
		name := req.FormValue("name")
		age, err := strconv.Atoi(req.FormValue("age"))
		if err != nil {
			fmt.Fprintf(w, "Sorry, not accepted age!")
			return
		}

		// Greet the user by name and age
		fmt.Fprintf(w, "Hello, %v year old named %s!", age, name)
	} else {
		// The request method is not POST, so redirect to the GET path
		http.Redirect(w, req, "/hello/", 302)
	}
}
