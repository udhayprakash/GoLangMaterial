package main

import (
	"flag"
	"fmt"
)

/*
go run f-parse-flags.go -flagvar=21 -flagvar2=AppDividend -flagvar3=false

Usage:
	-flag
	-flag=x
	-flag x
*/
func main() {
	var flagvar int
	var flagvar2 string
	var flagvar3 bool

	flag.IntVar(&flagvar, "flagvar", 111921, "Mandalorian Episode 4")
	flag.StringVar(&flagvar2, "flagvar2", "Gina Carano", "Mandalorian Episode")
	flag.BoolVar(&flagvar3, "flagvar3", true, "Mandalorian")

	flag.Parse()

	fmt.Println("flagvar :", flagvar)
	fmt.Println("flagvar2:", flagvar2)
	fmt.Println("flagvar3:", flagvar3)
}
