package main

import (
	"fmt"
	"reflect"
)

// Go does not provide a set type, but since the keys of a map
// are distinct, a map can serve this purpose

func main() {
	//uniqueValues := map[string]int{
	//	"red": 1,
	//	"white": 2,
	//	"red": 3,
	//}  // Compilation Error: duplicate key "red" in map literal

	colors := [...]string{"red", "white", "red", "blue", "yellow", "white", "red"}
	fmt.Println("colors                                  =", colors)

	uniqueColors := make(map[string]int)
	for _, color := range colors {
		_, alreadyPresent := uniqueColors[color]
		if !alreadyPresent {
			// add only if not present
			uniqueColors[color] = 1
		}
	}
	fmt.Println("uniqueColors                            =", uniqueColors)
	fmt.Println("reflect.ValueOf(uniqueColors).MapKeys() =", reflect.ValueOf(uniqueColors).MapKeys())
}
