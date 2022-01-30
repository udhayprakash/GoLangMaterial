package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Args[1:]:", os.Args[1:])
	for _, val := range os.Args[1:] {
		fmt.Printf("\t%T - %#[1]v\n", val)
	}
	fmt.Println()

	// This statement will parse the flag values
	// flag.Parse()

	fmt.Println("flag.Args():", flag.Args())

	for _, val := range flag.Args() {
		fmt.Printf("\t%T - %#[1]v\n", val)
	}
}

/*
OUTPUT:
--------
	~go run a-flag.go
	os.Args[1:]: []

	flag.Args(): []

	~go run a-flag.go 10
	os.Args[1:]: [10]
			string - "10"

	flag.Args(): [10]
			string - "10"
	~go run a-flag.go 10 20.2 true nil golang
	os.Args[1:]: [10 20.2 true nil golang]
			string - "10"        string - "20.2"
			string - "true"
			string - "nil"
			string - "golang"

	flag.Args(): [10 20.2 true nil golang]
			string - "10"
			string - "20.2"
			string - "true"
			string - "nil"
			string - "golang"


	~go run a-flag.go -h
		os.Args[1:]: [-h]
				string - "-h"

		flag.Args(): []
*/
