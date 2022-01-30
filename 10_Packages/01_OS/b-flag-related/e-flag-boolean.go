package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var toUpper bool
	flag.BoolVar(&toUpper, "u", false, "Should the words be converted to upper case")

	flag.Parse()

	words := flag.Args() // Even boolean flag value, also comes here

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
-------
    ~go run e-flag-boolean.go
    ~go run e-flag-boolean.go -h
    Usage of C:\Users\Amma\AppData\Local\Temp\go-build4183878898\b001\exe\e-flag-boolean.exe:  -u    Should the words be converted to upper case
    ~go run e-flag-boolean.go -u

    ~go run e-flag-boolean.go -u true
    TRUE

    ~go run e-flag-boolean.go -u true apple mango goa banana
    TRUEAPPLE
    MANGO
    GOA
    BANANA

*/
