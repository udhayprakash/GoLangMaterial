package main

import (
	"fmt"
	"time"
)

func main() {
	input := "2017-08-31"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	fmt.Println(t)                       // 2017-08-31 00:00:00 +0000 UTC
	fmt.Println(t.Format("02-Jan-2006")) // 31-Aug-2017
}
