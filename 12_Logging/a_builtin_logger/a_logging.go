package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("fmt - Hello world")
	log.Println("log - Hello world")

	fmt.Printf("fmt - Hello %s\n", "world")
	log.Printf("log - Hello %s\n", "world")
}
