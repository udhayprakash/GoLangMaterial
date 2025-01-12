package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://www.example.com:8000/user?username=Johnny&Gender=Male&age=21"

	// Parsing the URL content
	result, _ := url.Parse(s)

	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // www.example.com:8000
	fmt.Println(result.Path)     // /user
	fmt.Println(result.Port())   // 8000
	fmt.Println(result.RawQuery) // username=Johnny&Gender=Male&age=21

	// Extracting query values
	vals := result.Query()
	fmt.Println(vals["username"]) // [Johnny]
	fmt.Println(vals["Gender"])   // [Male]
	fmt.Println(vals["age"])      // [21]
}
