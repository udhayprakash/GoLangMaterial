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

	switch weekOfDay {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("Timings: 9 AM to 6 PM ")
	case "saturday":
		fmt.Println("Timings: 9 AM to 1 PM ")
	case "sunday":
		fmt.Println("----HOLIDAY ----------")
	default:
		fmt.Println("INVALID ENTRY")
	}

}
