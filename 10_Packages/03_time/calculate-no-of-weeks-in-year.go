package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
)

func main() {
	tn := time.Now() // change this to calculate different year

	fmt.Println("Now : ", tn.Format(time.ANSIC))

	// this week is which number in the current calender year
	_, currentWeekNumber := tn.ISOWeek()
	fmt.Println("Current week number : ", currentWeekNumber)

	// calculate the last week number for the current year
	last := now.EndOfYear()

	lastWeekOfTheYear := now.New(last).BeginningOfWeek()

	fmt.Println("The day of the final week : ", lastWeekOfTheYear.Format(time.ANSIC))
	_, lastWeekNumber := lastWeekOfTheYear.ISOWeek()

	// verify at https://www.epochconverter.com/weeks/2019
	fmt.Println("For the year", tn.Year(), ". It has", lastWeekNumber, "weeks.")

	weeksLeft := lastWeekNumber - currentWeekNumber

	fmt.Println("Hence, there are", weeksLeft, "weeks left to go.")
}
