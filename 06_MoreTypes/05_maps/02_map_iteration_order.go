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
	var keys []int
	for k := range myMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", myMap[k])
	}

	fmt.Println("\nAlternatively ...")
	for key, val := range myMap {
		fmt.Printf("\t %d : %s \n", key, val)
	}

	alphabets := map[int]string{
		64: "A",
		65: "B",
		66: "C",
	}
	fmt.Println(alphabets)

	// No builtin way to append maps
	for k2, v2 := range alphabets {
		myMap[k2] = v2
	}
	fmt.Println(myMap)

}
