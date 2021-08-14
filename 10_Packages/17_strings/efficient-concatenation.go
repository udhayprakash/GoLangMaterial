package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func main() {
	// Method 1
	var strs []string

	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		strs = append(strs,
			fmt.Sprintf("%d", i%10),
		)
	}
	fmt.Println(strings.Join(strs, ""))
	fmt.Println("TIME TAKEN:", time.Now().Sub(startTime))

	// Method 2
	var b bytes.Buffer

	startTime = time.Now()
	for i := 0; i < 1000; i++ {
		b.WriteString(
			fmt.Sprintf("%d", i%10),
		)
	}
	fmt.Println(b.String())
	fmt.Println("TIME TAKEN:", time.Now().Sub(startTime))
}
