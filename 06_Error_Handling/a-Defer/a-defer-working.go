package main

import "fmt"

func main() {
	defer fmt.Println("First")
	fmt.Println("Second")
	defer fmt.Println("Third")
}

// Defers will execute in LIFO (Last-In First-OUt) order

/*
 OUTPUT:

 	Second
	Third
	First
*/
