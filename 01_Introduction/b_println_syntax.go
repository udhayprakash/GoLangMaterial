package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	fmt.Println("Hello" + "world")
	fmt.Println("Hello", "world")
	fmt.Println("Hello", "world", 123, -2.2, true, nil)

	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	print("Hello world\n")
	println("Hello world")
}
