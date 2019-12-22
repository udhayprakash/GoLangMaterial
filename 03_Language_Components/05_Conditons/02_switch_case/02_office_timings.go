package main

import "fmt"

func main() {
	/*
		Monday 		9 AM to 6 PM
		Tuesday		9 AM to 6 PM
		Wednesday	9 AM to 6 PM
		Thursday	9 AM to 6 PM
		Friday 		9 AM to 6 PM
		Saturday	9 AM to 1 PM
		Sunday 		HOLIDAY
	*/
	fmt.Println("Enter day of the week:")

	var week_of_day string
	fmt.Scanf("%s", &week_of_day)

	switch week_of_day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("Timings: 9 AM to 6 PM")
	case "saturday":
		fmt.Println("Timings: 9 AM to 1 PM")
	case "sunday":
		fmt.Println("___ HOILDAY ____")
	default:
		fmt.Println("Invalid Entry")
	}
}
