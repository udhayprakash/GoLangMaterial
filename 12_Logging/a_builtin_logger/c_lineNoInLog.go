package main

import (
	"log"
)

func main() {

	// Enable line numbers in logging. By default, it is disabled
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Will print: "[date] [time] [filename]:[line]: [text]"
	log.Println("main - started")

	num1 := 100
	num2 := 200
	log.Println("num1 + num2:", num1+num2)

	log.Println("addition(num1, num2)=", addition(num1, num2))
	log.Println("main - end")
}

func addition(n1, n2 int) int {
	log.Println("addition function - start")
	return n1 + n2
}

// Assignment:
// If the log fortmatting flags were set in user-defined function, will it
// affect the formmting of logs from main() function