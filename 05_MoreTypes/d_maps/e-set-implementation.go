package main

import (
	"fmt"
	"reflect"
)

// Go does not provide a set type, but since the keys of a map
// are distinct, a map can serve this purpose

func main() {
	// uniqueValues := map[string]int{
	// 	"red": 1,
	// 	"white": 2,
	// 	"red": 3,
	// }  // Compilation Error: duplicate key "red" in map literal

	colors := [...]string{"red", "white", "red", "blue", "yellow", "white", "red"}
	fmt.Println("colors                                  =", colors)

	uniqueColors := make(map[string]int)
	// zero value of map 
	fmt.Println("uniqueColors == nil:", uniqueColors == nil)

	for _, color := range colors {
		// Lookup
		_, isKeyPresent := uniqueColors[color]
		if isKeyPresent == false {
			uniqueColors[color] = 1
		} else {
			uniqueColors[color] += 1
		}

	}
	fmt.Println("\nuniqueColors                            =", uniqueColors)
	fmt.Println("reflect.ValueOf(uniqueColors).MapKeys() =", reflect.ValueOf(uniqueColors).MapKeys())

	// To get non-duplicate colors
	var nonDuplicateColors []string
	for color, count := range uniqueColors {
		if count == 1 {
			nonDuplicateColors = append(nonDuplicateColors, color)
		}
	}
	fmt.Println("nonDuplicateColors=", nonDuplicateColors)
}
