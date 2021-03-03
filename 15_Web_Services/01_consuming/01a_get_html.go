package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://golangcode.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP Status Code and Status Name
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")
	} else {
		fmt.Println("Argh! Broken")
	}
}
