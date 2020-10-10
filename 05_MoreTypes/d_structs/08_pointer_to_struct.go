package main

import "fmt"

/*
Purpose: Creating a struct and obtaining a pointer to it
         using the built-in new() function
*/
type Car struct {
	chaseNumber int64
	model 	string
	launchYear int16
}

func main() {
	// You can also get a pointer to a struct using the built-in new() function
	// It allocates enough memory to fit a value of the given struct type, and returns a pointer to it
	nano := new(Car)

	nano.chaseNumber = 132323213132312131
	nano.model = "Tata Nano xyz"
	nano.launchYear = 2019
	fmt.Println(nano)




}
