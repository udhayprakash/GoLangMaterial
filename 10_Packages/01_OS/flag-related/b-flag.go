package main

import (
	"flag"
	"fmt"
)

func main() {
	var flagvar int
	flag.IntVar(&flagvar, "flagvar", 111921, "Mandalorian Episode 4")
	fmt.Println("Number", flagvar)
}
