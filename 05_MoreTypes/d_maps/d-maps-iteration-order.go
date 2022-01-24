package main

import (
	"fmt"
	"sort"
)

func main() {
	myMap := map[int]string{
		1: "d",
		2: "c",
		3: "b",
		4: "a",
	}
	fmt.Println("myMap:", myMap) // map[1:d 2:c 3:b 4:a]

	// To get all keys in map
	var keys []int
	for key := range myMap {
		keys = append(keys, key)
	}
	fmt.Println("All Keys 	:", keys) // [2 3 4 1]
	sort.Ints(keys)
	fmt.Println("sorted keys:", keys) // [1 2 3 4]

	// To get all the values
	var values []string
	for _, value := range myMap {
		values = append(values, value)
	}
	fmt.Println("\nAll values	:", values)
	sort.Strings(values)
	fmt.Println("sorted values  :", values)

	for key, value := range myMap {
		fmt.Println(key, "====>", value)
	}
	fmt.Println()

	// iterating on sorted keys
	for _, eachKey := range keys {
		fmt.Println(eachKey, "====>", myMap[eachKey])
	}
	fmt.Println()

	myMap2 := map[int]string{
		99: "nine",
		77: "seven",
		23: "two",
	}
	fmt.Println("\nmyMap :", myMap)
	fmt.Println("myMap2:", myMap2)

	// append two maps - No builtin way as of now
	for k2, v2 := range myMap2 {
		myMap[k2] = v2
	}
	fmt.Println("\nmyMap :", myMap)
	fmt.Println("myMap2:", myMap2)
}

/*
NOTE: When iterating over a map with a range loop, 
	the iteration order is not specified and is not 
	guaranteed to be the same from one iteration to the next.

	when stable order is expected in iteration, sort keys
	and access value from them.
*/