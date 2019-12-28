package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "123"
	//v = "two"
	if _, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%q looks like a number.\n", v)
	}
}

// Alternatively, use scanner.Scanner (from text/scanner) in mode ScanInts
