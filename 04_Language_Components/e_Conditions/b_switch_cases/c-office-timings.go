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
	println("Enter day of the week:")

	var weekOfDay string
	fmt.Scanf("%s", &weekOfDay)

	// converting to lower case
	weekOfDay = strings.ToLower(weekOfDay)

	switch weekOfDay {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		println("Timings: 9 AM to 6 PM ")
	case "saturday":
		println("Timings: 9 AM to 1 PM ")
	case "sunday":
		println("----HOLIDAY ----------")
	default:
		println("INVALID ENTRY")
	}

}