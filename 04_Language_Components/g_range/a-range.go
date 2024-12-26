package main

import "fmt"

// range can be used to iterate(loop) over any iterable object
// iterables - strings, bytes, array, slice, map
func main() {
	name := "Go Language"

	// G	o	 	L	a	n	g	u	a	g	e
	// 0 	1 	2	3	4	5	6	7	8	9	10

	// Method 1 - traditional
	for i := 0; i < len(name); i++ {
		fmt.Printf(" index= %2d  character= %q\n", i, name[i])
	}
	fmt.Println()

	// Method 2 - using range
	for i := range name {
		fmt.Printf(" index= %2d  character= %q\n", i, name[i])
	}
	fmt.Println()

	for i, eachChr := range name {
		fmt.Printf(" index= %2d  character= %q\n", i, eachChr) //name[i])
	}
	fmt.Println()

	for _, eachChr := range name {
		fmt.Printf("character= %q\n", eachChr)
	}
	fmt.Println("----")

	//----- Examples
	for i := 2; i < 5; i++ {
		fmt.Printf("character= %q\n", name[i])
	}
	fmt.Println()

	for _, eachChr := range name[2:5] { // passing a string slice
		fmt.Printf("character= %q\n", eachChr)
	}
	fmt.Println()

	for i, eachChr := range name {
		if i >= 2 && i < 5 {
			fmt.Printf(" index= %2d  character= %q\n", i, eachChr) //name[i])
		}
	}

	for i, eachChr := range name {
		if !(i >= 2 && i < 5) {
			continue
		}
		fmt.Printf(" index= %2d  character= %q\n", i, eachChr) //name[i])

	}
}
