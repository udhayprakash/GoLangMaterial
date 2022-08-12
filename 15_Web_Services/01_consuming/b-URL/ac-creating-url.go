package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Creating a URL from components
	newurl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/args",
		RawQuery: "x=1&y=2",
	}
	s := newurl.String()
	fmt.Println(s) // https://www.example.com/args?x=1&y=2

	newurl.Host = "facebook.com"
	s = newurl.String()
	fmt.Println(s) // https://facebook.com/args?x=1&y=2

	// === creating a new values struct and encode it as a query string
	newVals := url.Values{}
	newVals.Add("x", "100")
	newVals.Add("y", "12.3")
	newVals.Add("z", "true")

	newurl.RawQuery = newVals.Encode()

	s = newurl.String()
	fmt.Println(s) // https://facebook.com/args?x=100&y=12.3&z=true

}
