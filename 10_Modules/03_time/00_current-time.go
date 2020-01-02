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
	fmt.Println("nsecs:", nsecs)
	fmt.Println("msecs:", msecs)

	var timestamp int64 = 0
	fmt.Println(time.Unix(timestamp, 0)) // 1970-01-01 05:30:00 +0530 IST

}
