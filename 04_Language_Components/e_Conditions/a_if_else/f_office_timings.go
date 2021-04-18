package main

import (
	"fmt"
	"strings"
)

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

	var weekOfDay string
	fmt.Scanf("%s", &weekOfDay)

	// Converting to lower case
	weekOfDay = strings.ToLower(weekOfDay)

	if weekOfDay == "monday" {
		fmt.Println("Timings: 9 AM to 6 PM ")
	} else if weekOfDay == "tuesday" {
		fmt.Println("Timings: 9 AM to 6 PM ")
	} else if weekOfDay == "wednesday" {
		fmt.Println("Timings: 9 AM to 6 PM ")
	} else if weekOfDay == "thursday" {
		fmt.Println("Timings: 9 AM to 6 PM ")
	} else if weekOfDay == "friday" {
		fmt.Println("Timings: 9 AM to 6 PM ")
	} else if weekOfDay == "saturday" {
		fmt.Println("Timings: 9 AM to 1 PM ")
	} else if weekOfDay == "sunday" {
		fmt.Println("----HOLIDAY ----------")
	} else {
		fmt.Println("INVALID ENTRY")
	}

}
