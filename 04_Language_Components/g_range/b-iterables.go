package main

import (
	"fmt"
)

// Iterable object - object that supports looping over its elements
func main() {
	// value := true
	// for index, val := range value {
	// 	fmt.Println(index, val)
	// } // cannot range over value (type bool)

	// value := 12321
	// for index, val := range value {
	// 	fmt.Println(index, val)
	// } // cannot range over value (type int)

	// value := 123.21
	// for index, val := range value {
	// 	fmt.Println(index, val)
	// } // cannot range over value (type float64)

	// value := nil
	// for index, val := range value {
	// 	fmt.Println(index, val)
	// } // use of untyped nil

	value := "32344234" // strings are iterable
	for index, val := range value {
		fmt.Println(index, val)
	}
	fmt.Println()

	// value2 := '3'
	// for index, val := range value2 {
	// 	fmt.Println(index, val)
	// } // cannot range over value2 (type rune)

	// Looping over an Array
	numsArray := [7]int{11, 22, 33, 44, 55, 66, 77}
	//            0   1   2    3   4   5   6
	for index, val := range numsArray {
		fmt.Println(index, val)
	}

	// Looping over an Slice
	numsSlice := []int{11, 22, 33, 44, 55, 66, 77}
	//            0   1   2    3   4   5   6
	for index, val := range numsSlice {
		fmt.Println(index, val)
	}

	// Looping over a map - both key and value
	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cat"}

	for range kvs {
		fmt.Println("-in loop")
	}

	// Looping over a map - only key
	for k := range kvs {
		fmt.Println("key:", k)
	}

	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

}
