package main

import (
	"flag"
	"fmt"
)

func main() {
	var flagvar string
	flag.StringVar(&flagvar, "flagvar", "Gina Carano", "Mandalorian Episode 4")
	fmt.Println("Name", flagvar)
}
