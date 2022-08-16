// Fetch prints the content from a url
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE:\n fileName.go <URL>")
		fmt.Println("Example:03_fetch.go https://www.facebook.com/")
		os.Exit(1)
	}
}

func main() {

	fmt.Println(os.Args)
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// Assignment: try writing the outputs in files with the same domain name
