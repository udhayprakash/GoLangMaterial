package main

import (
	"fmt"
	"time"
)

/*
Try time.Now() in Go PLayground
It will return 2009-11-10 23:00:00 +0000 UTC m=+0.000000001
Because, It is the time and date of Go Lang's birthday.
*/
func main() {
	now := time.Now()       // current local time
	secs := now.Unix()      // number of seconds since January 1, 1970 UTC
	nsecs := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC

	// No milliseconds API. So, we can compute manually
	msecs := nsecs / 1000000

	fmt.Println("now  :", now)
	fmt.Println("secs :", secs)
	fmt.Println("msecs:", msecs)
	fmt.Println("nsecs:", nsecs)

	var timestamp int64 = 0
	fmt.Println(time.Unix(timestamp, 0))  // 1970-01-01 05:30:00 +0530 IST
	fmt.Println(time.Unix(secs, 0))       // 2021-11-11 19:07:17 +0530 IST
	fmt.Println(time.Unix(423332232, 0))  // 1983-06-01 21:47:12 +0530 IST
	fmt.Println(time.Unix(2342334249, 0)) // 2044-03-23 13:54:09 +0530 IST

}
