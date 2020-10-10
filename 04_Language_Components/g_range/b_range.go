package main

import "fmt"

func main(){
	// Looping over an Array
	nums := []int{11, 22, 33, 44, 55, 66, 77}

	for loopIndex, num := range nums {
		fmt.Printf("At loop index %d, value %d is present\n", loopIndex, num)
	}
	fmt.Println()

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


	// Looping over a string
	for i, c := range "go language" {
		fmt.Println(i, c)
		// first value is the starting byte index of the rune
		// second the rune itself.
	}

}
