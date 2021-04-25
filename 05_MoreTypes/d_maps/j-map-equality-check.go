package main

import (
	"fmt"
	"reflect"
)

func main() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	map2 := map1

	// if map1 == map2 { } // map can only be compared to nil

	fmt.Println("reflect.DeepEqual(map1, map2):", reflect.DeepEqual(map1, map2))
}
