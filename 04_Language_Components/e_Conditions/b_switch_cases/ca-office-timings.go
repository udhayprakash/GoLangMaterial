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
	fmt.Print("Enter day of the week:")

	var weekOfDay string
	fmt.Scanf("%s", &weekOfDay) // will pick only the characters

	// converting to lower case
	weekOfDay = strings.ToLower(weekOfDay)

	if weekOfDay == "monday" || weekOfDay == "tuesday" ||
		weekOfDay == "wednesday" || weekOfDay == "thursday" ||
		weekOfDay == "friday" {
		println("Timings: 9 AM to 6 PM ")
	} else if weekOfDay == "saturday" {
		println("Timings: 9 AM to 1 PM ")
	} else if weekOfDay == "sunday" {
		println("----HOLIDAY ----------")
	} else {
		println("INVALID ENTRY")
	}

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

// NOTE: else and default blocks are not mandatory
