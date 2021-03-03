package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	// Specify our ExampleHandler and wrap the Logger around it.
	http.Handle("/", Logger(http.HandlerFunc(ExampleHandler), "WEB"))

	log.Println("** Service Started on Port 8080 **")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

// Logger records incoming requests and uses the log package to print http method,
// route url and duration to screen/logs.
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Save start time to calculate duration
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}
