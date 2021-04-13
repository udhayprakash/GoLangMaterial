package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	fmt.Println("Hello" + "world")
	fmt.Println("Hello", "world")
	fmt.Println("Hello", "world", 123, -3.1416, true, nil)

	// f - formatting
	fmt.Printf("Hello\n")
	fmt.Printf("Hello %d\n", 123)

	fmt.Fprint(os.Stdout, "Hello ", 99, "\n")

	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))
	fmt.Println(fmt.Sprintf("Hello %d", 23))

	print("Hello world\n")
	println("Hello world")
}
