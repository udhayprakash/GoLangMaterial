package main

import (
	"fmt"
	"time"
)

func clockDuration(clock string) (time.Duration, error) {
	c, err := time.Parse("15:04:05", clock)
	if err != nil {
		return 0, err
	}
	h, m, s := c.Clock()
	d := time.Duration(h)*time.Hour +
		time.Duration(m)*time.Minute +
		time.Duration(s)*time.Second
	return d, nil
}

func main() {
	d, err := clockDuration("01:02:03")
	fmt.Println(d, err)
}