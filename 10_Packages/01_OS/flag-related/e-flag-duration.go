package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()

}
/*

~go run e-flag-duration.go -period 10ms
	Sleeping for 10ms...

~go run e-flag-duration.go -period 3s  
	Sleeping for 3s...

go run e-flag-duration.go -period 10
	invalid value "10" for flag -period: parse errorUsage of C:\Users\Amma\AppData\Local\Temp\go-build3362939746\b001\exe\e-flag-duration.exe:  -period duration        sleep period (default 1s)exit status 2



*/