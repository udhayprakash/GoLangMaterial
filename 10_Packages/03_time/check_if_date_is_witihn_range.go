package main

import (
	"fmt"
	"time"
)

func DateWithinRange(dateString string, dateFormat string) (bool, error) {

	dateStamp, err := time.Parse(dateFormat, dateString)

	if err != nil {
		return false, err
	}

	today := time.Now()

	twoMonthsAgo := today.AddDate(0, -2, 0)  // minus 2 months
	twoMonthsLater := today.AddDate(0, 2, 0) // plus 2 months

	fmt.Println("\nGiven          		: ", dateStamp.Format("02/01/2006"))
	fmt.Println("2 months ago   		: ", twoMonthsAgo.Format("02/01/2006"))
	fmt.Println("2 months later 		: ", twoMonthsLater.Format("02/01/2006"))

	if dateStamp.Before(twoMonthsLater) && dateStamp.After(twoMonthsAgo) {
		return true, nil
	} else {
		return false, nil
	}

	// default
	return false, nil
}

func main() {

	// dd/mm/yyyy
	result, err := DateWithinRange("27/04/2017", "02/01/2006")

	if err != nil {
		fmt.Println("Error :", err)
	}

	// %t to print boolean value in fmt.Printf
	fmt.Printf("Given date with range 	: [%t]\n", result)

	// yyyy-dd-mm
	result1, err := DateWithinRange("2017-27-04", "2006-02-01")

	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Printf("Given date with range 	: [%t]\n", result1)

	// out of range example
	result2, err := DateWithinRange("2011-27-04", "2006-02-01")

	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Printf("Given date with range 	: [%t]\n", result2)

	// out of range example
	result3, err := DateWithinRange("2018-17-09", "2006-02-01")

	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Printf("Given date with range 	: [%t]\n", result3)
}
