package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()      // current local time
	sec := now.Unix()      // number of seconds since January 1, 1970 UTC
	nsec := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC

	fmt.Println(now)
	fmt.Println(sec)
	fmt.Println(nsec)

	var timestamp int64 = 0
	fmt.Println(time.Unix(timestamp, 0))  // 1970-01-01 05:30:00 +0530 IST

}
