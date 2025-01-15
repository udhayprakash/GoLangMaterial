package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	// Create a cookie which lasts 30 mins
	addCookie(w, "TestCookieName", "TestValue", 30*time.Minute)
	// Write our example
	io.WriteString(w, "Hello world!")
}

// addCookie will apply a new cookie to the response of a http request
// with the key/value specified.
func addCookie(w http.ResponseWriter, name, value string, ttl time.Duration) {
	expire := time.Now().Add(ttl)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}
