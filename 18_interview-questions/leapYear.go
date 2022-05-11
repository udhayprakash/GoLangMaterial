package main

import "fmt"

func main() {
	// leapYear(1998)
	// leapYear(2001)
	// leapYear(2004)
	for year := 2000; year <= 2022; year++ {
		leapYear(year)
	}
}

func leapYear(givenYear int) {
	if givenYear%4 == 0 {
		if givenYear%100 != 0 {
			fmt.Printf("%d is LEAP year\n", givenYear)
		} else {
			fmt.Printf("%d is NOT LEAP year\n", givenYear)
		}
	} else {
		fmt.Printf("%d is NOT LEAP year\n", givenYear)
	}
}
