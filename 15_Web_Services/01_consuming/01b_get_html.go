package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://golangcode.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// The line below would fail because Body = io.ReadCloser
	// fmt.Printf(response.Body)

	// ...so we convert it to a string by passing it through
	// a buffer first. A 'costly' but useful process.
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	fmt.Printf(newStr)
}
