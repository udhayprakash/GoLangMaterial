package main

// Command-line flags are a common way to specify options for command-line programs.
//For example, in wc -l the -l is a command-line flag.

import (
	"flag" // for flag parsing
	"fmt"
)

func main() {

	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args()) // similar to os.Args

}

/*
To get the usage
	go run h-example.go --help
	go run h-example.go -h

Usage
	go run h-example.go -word=opt -numb=7 -fork -svar=flag


When some args were not given, they will take default values
	go run h-example.go -word=opt

Trailing positional arguments can be provided after any flags.
	go run h-example.go -word=opt a1 a2 a3

Note that the flag package requires all flags to appear before positional arguments (otherwise the flags will be interpreted as positional arguments).
	go run h-example.go -word=opt a1 a2 a3 -numb=7
*/
