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
	flag.Parse()
	fmt.Println("flag.Args():", flag.Args())
	for _, val := range flag.Args() {
		fmt.Printf("\t%T - %#[1]v\n", val)
	}

}

/*

 $ go run a-flagless_values.go 
os.Args[1:]: []

$ go run a-flagless_values.go 213 123.123 True nil 
os.Args[1:]: [213 123.123 True nil]
        string - "213"
        string - "123.123"
        string - "True"
        string - "nil"

$ go run a-flagless_values.go 
os.Args[1:]: []

flag.Args(): []
$ go run a-flagless_values.go 213 123.123 True nil 
os.Args[1:]: [213 123.123 True nil]
        string - "213"
        string - "123.123"
        string - "True"
        string - "nil"

flag.Args(): []
$ go run a-flagless_values.go 213 123.123 True nil 
os.Args[1:]: [213 123.123 True nil]
        string - "213"
        string - "123.123"
        string - "True"
        string - "nil"

flag.Args(): [213 123.123 True nil]
$ go run a-flagless_values.go 213 123.123 True nil 
os.Args[1:]: [213 123.123 True nil]
        string - "213"
        string - "123.123"
        string - "True"
        string - "nil"

flag.Args(): [213 123.123 True nil]
        string - "213"
        string - "123.123"
        string - "True"
        string - "nil"
$ go run a-flagless_values.go 
os.Args[1:]: []

flag.Args(): []
$ go run a-flagless_values.go -h
os.Args[1:]: [-h]
        string - "-h"

Usage of /tmp/go-build3489782134/b001/exe/a-flagless_values:


*/
