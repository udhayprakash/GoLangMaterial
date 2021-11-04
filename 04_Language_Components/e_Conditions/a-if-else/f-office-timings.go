package main

import (
	"fmt"
	"strings"
)

/*
	OFFICE TIMINGS
		Monday 		9 AM to 6 PM
		Tuesday		9 AM to 6 PM
		Wednesday	9 AM to 6 PM
		Thursday	9 AM to 6 PM
		Friday 		9 AM to 6 PM
		Saturday	9 AM to 1 PM
		Sunday 		HOLIDAY
*/
func main() {
	fmt.Println("Enter day of the week:")

	var weekOfDay string
	fmt.Scanf("%s", &weekOfDay) // will pick only the characters

	// conversting to lower case
	weekOfDay = strings.ToLower(weekOfDay)

	if weekOfDay == "monday" {
		println("Timings: 9 AM to 6 PM ")
	} else if weekOfDay == "tuesday" {
		println("Timings: 9 AM to 6 PM ")
	} else if weekOfDay == "wednesday" {
		println("Timings: 9 AM to 6 PM ")
	} else if weekOfDay == "thursday" {
		println("Timings: 9 AM to 6 PM ")
	} else if weekOfDay == "friday" {
		println("Timings: 9 AM to 6 PM ")
	} else if weekOfDay == "saturday" {
		println("Timings: 9 AM to 1 PM ")
	} else if weekOfDay == "sunday" {
		println("----HOLIDAY ----------")
	} else {
		println("INVALID ENTRY")
	}
}

// NOTE: else block  is not mandatory
 