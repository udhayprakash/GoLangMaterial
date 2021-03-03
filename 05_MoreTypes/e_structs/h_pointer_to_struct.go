package main

import "fmt"

/*
Purpose: Creating a struct and obtaining a pointer to it
         using the built-in new() function
*/

type Car struct {
	chaseNumber int64
	model       string
	launchYear  int16
}

func main() {
	nano := Car{chaseNumber: 234234242432, model: "Tata Nano", launchYear: 2009}
	fmt.Println("nano = ", nano) // nano =  {234234242432 Tata Nano 2009}

	nano.chaseNumber = 3242342
	nano.model = "something"
	nano.launchYear = 2019
	fmt.Println("nano = ", nano) //nano =  {3242342 something 2019}

	// You can also get a pointer to a struct using the built-in new() function
	// It allocates enough memory to fit a value of the given struct type, and returns a pointer to it
	sumo := new(Car)

	sumo.chaseNumber = 234234
	sumo.model = "Tata Sumo"
	sumo.launchYear = 2006
	fmt.Println("sumo = ", sumo) // sumo =  &{234234 Tata Sumo 2006}

}
