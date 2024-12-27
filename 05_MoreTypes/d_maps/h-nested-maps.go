package main

import "fmt"

/*
// Map of Maps

	{
		"a": {
			"a1": 1,
			"a2": 2,
		},
		"b": {
			"b1": 1,
			"b2": 2,
		}
	}
*/
func main() {

	myMap1 := map[string]map[string]int{
		"a": map[string]int{
			"a1": 1,
			"a2": 2,
		},
		"b": map[string]int{
			"b1": 1,
			"b2": 2,
		},
	}
	fmt.Println("myMap1 = ", myMap1)
	//  map[a:map[a1:1 a2:2] b:map[b1:1 b2:2]]

	// map of arrays
	myMap2 := map[string][2]int{
		"one":   [2]int{1, 11},
		"two":   [...]int{2, 22},
		"three": {3, 33},
	}
	fmt.Println("myMap2 = ", myMap2)
	//  map[one:[1 11] three:[3 33] two:[2 22]]

	// map of slices
	myMap3 := map[string][]float64{
		"one":   []float64{1.1, 1.2, 1.3},
		"two":   []float64{2.0},
		"three": {3.1, 3.2, 3.3, 3.4, 3.5},
	}
	fmt.Println("myMap3 = ", myMap3)
	// map[one:[1.1 1.2 1.3] three:[3.1 3.2 3.3 3.4 3.5] two:[2]
}
