package main

import "fmt"

/*
	converts from feet into meters.
        (1 ft = 0.3048 m)
*/
func main() {

	var (
		meters float32
		feet   float32
	)
	fmt.Printf(`
		Initial default(zero) values:
			meters = %f
			feet   = %f
	`, meters, feet)

	// Get the meters value, in runtime
	fmt.Print("Enter no. of meters =")
	fmt.Scanf("%f", &meters)

	feet = 0.3048 * meters

	fmt.Printf(`
		finally:
			meters = %f
			feet   = %f`, meters, feet)

}
