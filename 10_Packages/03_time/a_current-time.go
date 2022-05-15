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

	fmt.Println("now  :", now)   // 2022-05-14 12:49:59.7929117 +0530 IST m=+0.005778101
	fmt.Println("secs :", secs)  // 1652512799
	fmt.Println("msecs:", msecs) // 1652512799792
	fmt.Println("nsecs:", nsecs) // 1652512799792911700

	var timestamp int64 = 0
	fmt.Println(time.Unix(timestamp, 0))  // 1970-01-01 05:30:00 +0530 IST
	fmt.Println(time.Unix(secs, 0))       // 2022-05-14 12:49:59 +0530 IST
	fmt.Println(time.Unix(423332232, 0))  // 1983-06-01 21:47:12 +0530 IST
	fmt.Println(time.Unix(2342334249, 0)) // 2044-03-23 13:54:09 +0530 IST

}
