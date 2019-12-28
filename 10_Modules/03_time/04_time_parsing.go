package main

import (
	"fmt"
	"time"
)

/*
Purpose: Parsing time
	Here, to parse a date you don’t specify the date using letters
	(Y-m-d for example). Instead you use a real time as the format
	- this time in fact: 2006-01-02T15:04:05+07:00.

	----------------------------------------------------
	standard		Format
	----------------------------------------------------
	RFC822			02 Jan 06 15:04 MST
	RFC850			Monday, 02-Jan-06 15:04:05 MST
	RFC1123			Mon, 02 Jan 2006 15:04:05 MST
	RFC3339			2006-01-02T15:04:05Z07:00
	RFC3339Nano		2006-01-02T15:04:05.999999999Z07:00

To use the MySQL date format, you’d use: “2006-01-02 15:04:05”

*/

func main() {
	input := "2017-08-31"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	fmt.Println(t)                       // 2017-08-31 00:00:00 +0000 UTC
	fmt.Println(t.Format("02-Jan-2006")) // 31-Aug-2017
	fmt.Println()

	input2 := "2018-01-20 04:35"
	layout2 := "2006-01-02 15:04"
	myDate, err := time.Parse(layout2, input2)
	if err != nil {
		panic(err)
	}

	// Format uses the same formatting style as parse, or we can use a pre-made constant
	fmt.Println("My Date Reformatted\t:", myDate.Format(time.RFC822)) // 20 Jan 18 04:35 UTC

	// In Y-m-d
	fmt.Println("Just The Date\t\t:", myDate.Format("2006-01-02")) // 2018-01-20
}
