package main

import "fmt"

// Go does not provide a set type, but since the keys of a map
// are distinct, a map can serve this purpose

func main() {
	// uniqueValues := map[string]int{
	// 	"red": 1,
	// 	"white": 2,
	// 	"red": 3,
	// }  // Compilation Error: duplicate key "red" in map literal

	colors := [...]string{"red", "white", "red", "blue", "yellow", "white", "red"}
	fmt.Println("colors       =", colors)

	uniqueColors := make(map[string]int)
	// zero value of map
	fmt.Println("uniqueColors == nil:", uniqueColors == nil)

	// Method 1
	for _, color := range colors {
		// fmt.Println(color)

		// Lookup
		_, isKeyPresent := uniqueColors[color]
		if isKeyPresent == false {
			uniqueColors[color] = 1
		} else {
			uniqueColors[color]++
		}
	}
	fmt.Println("\nuniqueColors    =", uniqueColors)

	// Method 2
	uniqueColors = make(map[string]int)
	for _, color := range colors {
		// Lookup
		val, _ := uniqueColors[color]
		uniqueColors[color] = val + 1
	}
	fmt.Println("\nuniqueColors    =", uniqueColors)
	fmt.Println("reflect.ValueOf(uniqueColors).MapKeys() =", reflect.ValueOf(uniqueColors).MapKeys())

	// To get non-duplicate colors
	var nonDuplicateColors []string
	for color, count := range uniqueColors {
		if count == 1 {
			// fmt.Println(color)
			nonDuplicateColors = append(nonDuplicateColors, color)
		}
	}
	fmt.Println("nonDuplicateColors=", nonDuplicateColors)

}

// Assignment : Try to get all colors, which are occurring more than 2 times
