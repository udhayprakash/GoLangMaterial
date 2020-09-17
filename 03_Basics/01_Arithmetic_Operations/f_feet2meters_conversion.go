/*
	converts from feet into meters.
        (1 ft = 0.3048 m)
*/

package main

import "fmt"

func main() {
	//const (
	//	meters float32
	//	feet float32
	//) // const declaration cannot have type without expression

	var (
		meters float32
		feet   float32
	)
	fmt.Println("Initial default Values:")
	fmt.Println("\tmeters =", meters)
	fmt.Println("\tfeet   =", feet)
	fmt.Println()

	// Get the meters value in run time
	fmt.Print("Enter no. of meters =")
	fmt.Scanf("%f", &meters)
	feet = 0.3048 * meters

	fmt.Println("\tmeters =", meters)
	fmt.Println("\tfeet   =", feet)
}
