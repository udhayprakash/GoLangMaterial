package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var toUpper bool
	flag.BoolVar(&toUpper, "u", false, "should the wordls be converted to upper case")

	flag.Parse()

	words := flag.Args()

	// // Method 1 - expensive operation
	// for _, word := range words {
	// 	if toUpper == true {
	// 		fmt.Println(strings.ToUpper(word))
	// 	} else {
	// 		fmt.Println(word)
	// 	}

	// }

	// Method 2 - optimal
	if toUpper == true {
		for _, word := range words {
			fmt.Println(strings.ToUpper(word))
		}
	} else {
		for _, word := range words {
			fmt.Println(word)
		}
	}

}


/*
OUTPUT:
------
$ go run c-flag-boolean.go 
$ go run c-flag-boolean.go -h
Usage of /tmp/go-build3541989184/b001/exe/c-flag-boolean:
  -u    should the wordls be converted to upper case
$ go run c-flag-boolean.go -u
$ go run c-flag-boolean.go -u=true
$ go run c-flag-boolean.go -u=true apple mango goa banana
APPLE
MANGO
GOA
BANANA
*/
