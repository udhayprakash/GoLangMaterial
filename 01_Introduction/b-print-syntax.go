package main

import (
	"fmt"
	"os"
)

func main() {
	// println
	fmt.Println("Hello")
	fmt.Println("Hello" + "World")
	fmt.Println("Hello", "World")
	fmt.Println("Hello", "World", 123, -3.1416, true, nil)

	fmt.Println() // Empty line

	// print
	fmt.Print("Hello\n")
	fmt.Print("Hello" + "World\n")
	fmt.Print("Hello", "World\n")
	fmt.Print("Hello", "World", 123, -3.1416, true, nil, "\n")

	fmt.Println() // Empty line

	// printf: f - formatting
	fmt.Printf("Hello\n")
	fmt.Printf("Hello" + "World\n")
	fmt.Printf("Hello", "World\n")
	fmt.Printf("Hello", "World", 123, -3.1416, true, nil, "\n")

	fmt.Printf("Hello %d\n", 123)
	fmt.Printf("Hello %s\n", "Golang")
	fmt.Printf("Name= %s  Age=%d salary=%f\n", "Edson", 34, 3432.4)

	fmt.Fprint(os.Stdout, "Hello ", 99, "\n")

	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))
	fmt.Println(fmt.Sprintf("Hello %d", 23))

	print("Hello world\n")
	println("Hello world")
}

// Ref: https://cheatography.com/fenistil/cheat-sheets/go-fmt-formattings/
