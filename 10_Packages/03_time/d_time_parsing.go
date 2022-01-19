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
const longFormLayout = "Jan 2, 2006 at 3:04pm (MST)"
const shortFormLayout = "2006-Jan-02"

func main() {

	// Parsing time units string
	m, _ := time.ParseDuration("1m30s")
	fmt.Printf("Take off in t-%.0f seconds.\n", m.Seconds())

	input := "2020-08-31"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	fmt.Println(t)                       // 2020-08-31 00:00:00 +0000 UTC
	fmt.Println(t.Format("02-Jan-2006")) // 31-Aug-2020
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

	// ISO 8601 Time
	fmt.Println("ISO 8601 Time\t\t:", myDate.Format(time.RFC3339)) // 20 Jan 18 04:35 UTC

	t2, _ := time.Parse(longFormLayout, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t2) // 2013-02-03 19:54:00 +0000 PST

	t3, _ := time.Parse(shortFormLayout, "2013-Feb-03")
	fmt.Println(t3) // 2013-02-03 00:00:00 +0000 UTC

	// Parse Unix Date format
	t4, err := time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}
	fmt.Println("default format:", t4)                    // 2015-02-25 11:06:39 +0000 PST
	fmt.Println("Unix format:", t4.Format(time.UnixDate)) // Wed Feb 25 11:06:39 PST 2015

	/*Pre-defined Layouts
	  ------------------------------------------------------
	  Label				Definition
	  ------------------------------------------------------
	  Layout				01/02 03:04:05PM '06 -0700
	  ANSIC				Mon Jan _2 15:04:05 2006
	  UnixDate			Mon Jan _2 15:04:05 MST 2006
	  RubyDate			Mon Jan 02 15:04:05 -0700 2006
	  RFC822				02 Jan 06 15:04 MST
	  RFC822Z				02 Jan 06 15:04 -0700
	  RFC850				Monday, 02-Jan-06 15:04:05 MST
	  RFC1123				Mon, 02 Jan 2006 15:04:05 MST
	  RFC1123Z			Mon, 02 Jan 2006 15:04:05 -0700
	  RFC3339				2006-01-02T15:04:05Z07:00
	  RFC3339Nano			2006-01-02T15:04:05.999999999Z07:00
	  Kitchen				3:04PM
	  Stamp				Jan _2 15:04:05
	  StampMilli			Jan _2 15:04:05.000
	  StampMicro			Jan _2 15:04:05.000000
	  StampNano			Jan _2 15:04:05.000000000
	  ------------------------------------------------------

	  ================
	  formatting Rules
	  ================
		----------------------------------------------------------
		Type							Options
		----------------------------------------------------------
		Year							"2006", "06"
		Month							"Jan", "January", "1", "01"
		Textual day of the week			"Mon", "Monday"
		Numeric day of the month		"2", "_2", "02"
		Numeric day of the year			"__2", "002"
		Hour							"15", "3", "03" (PM or AM)
		Minute							"4", "04"
		Second							"5", "05"
		AM/PM mark						"PM"
		----------------------------------------------------------

		For layouts specifying the two-digit year 06, a value NN >= 69 will
		be treated as 19NN and a value NN < 69 will be treated as 20NN.

		----------------------------------
		Value			Format
		----------------------------------
		"-0700"			±hhmm
		"-07:00"		±hh:mm
		"-07"			±hh
		"Z0700"			Z or ±hhmm
		"Z07:00"		Z or ±hh:mm
		"Z07"			Z or ±hh
		----------------------------------

	*/
	ipTimes := []string{
		"24 Aug 2021 7:20:50 PM",
		"Aug 24 2021 07:20:50PM",
		"24-08-2021 17:20:50",
		"08-24-2021 13:12:32",
		"24/8/21 9:00:12",
		"8/24/2021 10:10:10",
		"1993-11-23 7**12**44PM",
		"24 October 97 7:12:45PM+05:30", // 24 th January 1997
	}

	layouts := []string{
		"2 Jan 2006 3:04:05 PM",
		"Jan 2 2006 03:04:05PM",
		"2-01-2006 15:04:05",
		"01-02-2006 15:4:5",
		"2/1/06 3:4:5",
		"1/2/2006 3:4:5",
		"2006-01-02 3**4**5PM",
		"2 January 06 3:4:5PMZ07:00",
	}

	for i := 0; i < len(ipTimes); i++ {
		t, err := time.Parse(layouts[i], ipTimes[i]) // parsing to get time value
		if err != nil {                              // check for errors
			fmt.Println(err)
			return
		}

		fmt.Println("    Input DateTime String: ", ipTimes[i])
		fmt.Println("                   Layout: ", layouts[i])
		fmt.Println("Time value Default Format: ", t)
		fmt.Println("   Time in RFC3339 Format: ", t.Format(time.RFC3339))
	}
}
