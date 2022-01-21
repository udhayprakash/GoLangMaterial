package main

import (
	"fmt"
	"time"
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

	fmt.Println("time.Now().Weekday()=", time.Now().Weekday())

	switch time.Now().Weekday() { // switch initilizer

	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Timings: 9 AM to 6 PM ")
	case time.Saturday:
		fmt.Println("Timings: 9 AM to 1 PM ")
	case time.Sunday:
		fmt.Println("----HOLIDAY ----------")
	default:
		fmt.Println("INVALID ENTRY")
	}

}
