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
	fmt.Println("myMap:", myMap)

	var keys []int
	for k := range myMap {
		//fmt.Println("append(keys, k):", append(keys, k))
		keys = append(keys, k)
	}
	fmt.Println("Keys:", keys)

	sort.Ints(keys)
	fmt.Println("After sort.Ints(keys), keys:", keys)

	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", myMap[k])
	}
	fmt.Println()

	fmt.Println("\nAlternatively ...")
	for key, val := range myMap {
		fmt.Printf("\t %d : %s \n", key, val)
	}

	alphabets := map[int]string{
		64: "A",
		65: "B",
		66: "C",
	}
	fmt.Println("alphabets:", alphabets)

	// No builtin way to append maps
	for k2, v2 := range alphabets {
		myMap[k2] = v2
	}
	fmt.Println("myMap:", myMap)

}
