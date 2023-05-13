package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	resp, err := http.Get("https://golang.org/")
	CheckError(err)

	defer resp.Body.Close()

	// Print the HTTP Status Code and Status Name
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	// The line below would fail because Body = io.ReadCloser
	// fmt.Printf(response.Body)

	// ...so we convert it to a string by passing it through
	// a buffer first. A 'costly' but useful process.
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	newStr := buf.String()
	fmt.Printf(newStr)

}
