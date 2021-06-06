package main

import (
	"flag"
	"fmt"
)

func main() {
	ip := flag.Int("num", 111921, "Mandalorian Episode 4")
	fmt.Println("Number", *ip)
}
