package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
Middleware is an entity that intercepts the server's request/response life cycle.
  - Process the request before running business logic (authentication)
  - Modify the request to the next handler function (attaching payload)
  - Modify the response for the client
  - Logging.... and much more

*/
func main() {
	http.HandleFunc("/hello", timed(hello))
	fmt.Println("Listening at localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		time.Sleep(2)
		end := time.Now()
		fmt.Println("The request took", end.Sub(start))
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}
