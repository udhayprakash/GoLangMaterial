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
}
