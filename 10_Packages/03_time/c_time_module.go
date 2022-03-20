package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world")

	p := fmt.Println // Aliasing a function
	p("Hello world")

	now := time.Now()
	p(now) // 2021-11-12 13:09:51.8732094 +0530 IST m=+0.004198701

	then := time.Date(
		1947, 11, 12, 20, 34, 58, 651387237, time.UTC)
	p(then) // 1947-11-12 20:34:58.651387237 +0000 UTC

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday()) // Wednesday

	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	diff := now.Sub(then)
	p(diff) // 648683h8m15.184136763s

	p(diff.Hours())       // 648683.1375511491
	p(diff.Minutes())     // 3.892098825306895e+07
	p(diff.Seconds())     // 2.335259295184137e+09
	p(diff.Nanoseconds()) // 2335259295184136763

	p(then.Add(diff))  // 2021-11-12 07:43:13.835524 +0000 UTC
	p(then.Add(-diff)) // 1873-11-11 09:26:43.467250474 +0000 UTC

}
