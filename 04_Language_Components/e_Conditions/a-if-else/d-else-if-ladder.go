package main

import "fmt"

/*
	+90	North Pole
		Northern Hemisphere
	0 	Equator
		Southern Hemisphere
	-90 South Pole

*/
func main() {
	var latitude float32
	fmt.Println("Please enter the latitude degrees:")
	fmt.Scanf("%f", &latitude)

	fmt.Println("You entered:", latitude)
	if latitude == 90 {
		fmt.Println("You are located at North Pole")
	} else if latitude < 90 && latitude > 0 {
		fmt.Println("you are located in North Hemisphere")
	} else if latitude == 0 {
		fmt.Println("YOu are located on the EQUATOR")
	} else if latitude < 0 && latitude > -90 {
		fmt.Println("You are located in Southern Hemisphere")
	} else if latitude == -90 {
		fmt.Println("YOu are located at South Pole")
	} else {
		fmt.Println("You Entered invalid latitude degrees!!!")
	}

}

// NOTE: else block is optional
