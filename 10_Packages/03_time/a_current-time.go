package main

import (
	"fmt"
	"time"
)

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
	fmt.Println(time.Unix(2342334249, 0)) // 2044-03-23 13:54:09 +0530 IST

}
