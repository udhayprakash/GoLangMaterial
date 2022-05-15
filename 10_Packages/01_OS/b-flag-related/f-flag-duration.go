package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var period = flag.Duration("period", 1*time.Second, "sleep period")

	flag.Parse()

	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)

	fmt.Println("Done")
}

/*
OUTPUT:
--------
	~go run f-flag-duration.go -period 10
	invalid value "10" for flag -period: parse error
	Usage of C:\..Temp\go-build..\exe\f-flag-duration.exe:
	-period duration
			sleep period (default 1s)
	exit status 2

	~go run f-flag-duration.go -period 10ms
		Sleeping for 10ms...

	~go run f-flag-duration.go -period 3s
		Sleeping for 3s...

	go run f-flag-duration.go -period 10
		invalid value "10" for flag -period: parse error
	Usage of C:\...\Temp\go-build...\exe\e-flag-duration.exe:  -period duration        sleep period (default 1s)
	exit status 2
*/
