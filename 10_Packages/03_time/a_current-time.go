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

	fmt.Println("now  :", now)   // 2025-01-01 06:31:41.661112861 +0000 UTC m=+0.000028534
	fmt.Println("secs :", secs)  // 1735713343
	fmt.Println("msecs:", msecs) // 1735713343484
	fmt.Println("nsecs:", nsecs) // 1735713343484236401

	var timestamp int64 = 0
	fmt.Println(time.Unix(timestamp, 0))  // 1970-01-01 00:00:00 +0000 UTC
	fmt.Println(time.Unix(secs, 0))       // 2025-01-01 06:37:08 +0000 UTC
	fmt.Println(time.Unix(31212, 0))      // 1970-01-01 08:40:12 +0000 UTC
	fmt.Println(time.Unix(2342334249, 0)) // 2044-03-23 08:24:09 +0000 UTC
}
