package main

import (
	"fmt"
	"time"
)

const clockLayout = "15:04:05"

var clockZero, _ = time.Parse(clockLayout, "00:00:00")

func ClockDuration(clock string) (time.Duration, error) {
	c, err := time.Parse(clockLayout, clock)
	if err != nil {
		return 0, err
	}
	return c.Sub(clockZero), nil
}

func main() {
	d, err := ClockDuration("01:02:03")
	fmt.Println(d, err)
}
