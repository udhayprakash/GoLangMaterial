package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	originalURL := "http://google.com"
	originalURL = "http://geocities.com"
	originalURL = "//socketloop.com"  // error
	originalURL = "socketloop.com"  // error
	originalURL = "https://socketloop.com" 

	resp, err := http.Get(originalURL)
	if err != nil {
		fmt.Println(err)
	}

	finalURL := resp.Request.URL.String()

	fmt.Println("Original URL is : ", originalURL)
	fmt.Println("Final URL is    : ", finalURL)

	// Check if served with https
	fmt.Println("Is HTTPS ? : ", strings.HasPrefix(finalURL, "https"))
}
