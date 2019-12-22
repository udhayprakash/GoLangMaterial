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

	if week_of_day == "monday" {
		fmt.Println("Timings: 9 AM to 6 PM")
	} else if week_of_day == "tuesday" {
		fmt.Println("Timings: 9 AM to 6 PM")
	} else if week_of_day == "wednesday" {
		fmt.Println("Timings: 9 AM to 6 PM")
	} else if week_of_day == "thursday" {
		fmt.Println("Timings: 9 AM to 6 PM")
	} else if week_of_day == "friday" {
		fmt.Println("Timings: 9 AM to 6 PM")
	} else if week_of_day == "saturday" {
		fmt.Println("Timings: 9 AM to 1 PM")
	} else if week_of_day == "sunday" {
		fmt.Println("___ HOILDAY ____")
	} else {
		fmt.Println("Invalid Entry")
	}
}

