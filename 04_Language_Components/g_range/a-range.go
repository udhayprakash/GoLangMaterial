package main

import (
	"fmt"
)

// range can be used to iterate(loop) over any iterable object
// iterables -
func main() {
	name := "Go Language"

	// G	o	 	L	a	n	g	u	a	g	e
	// 0 	1 	2	3	4	5	6	7	8	9	10

	// Method 1
	for i := 0; i < len(name); i++ {
		fmt.Printf("%q\n", name[i])
	}

	// Method 2
	for index := range name {
		fmt.Println("index=", index)
	}

	for index, each_chr := range name {
		fmt.Printf("index= %d\t each_chr=%q\n", index, each_chr)
	}
}
