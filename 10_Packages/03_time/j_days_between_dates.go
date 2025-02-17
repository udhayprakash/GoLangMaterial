package main

import (
	"fmt"
	"time"
)

func main() {
	// The leap year 2000 had 366 days.
	t1 := Date(2000, 1, 1)
	t2 := Date(2001, 1, 1)
	days := t2.Sub(t1).Hours() / 24
	fmt.Println("days :", days)

}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}


// Assignment - Patch calendar design for linux and windows systems
//        linux patch life is 1 per Quarter on first saturday
// 		  window patch life is 1 per month on first saturday
// for a given year
