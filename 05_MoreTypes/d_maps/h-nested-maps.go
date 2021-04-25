package main

import (
	"fmt"
)

func main (){
	myMap1 := map[string]map[string]string{
		"one": map[string]string{
			"one-a": "1A",
			"one-b": "1B",
		},
		"two": {
			"two-a": "2A",
		},
	}

	fmt.Println("myMap1 = ", myMap1)

	// map of arrays
	myMap2 := map[string][2]int{
		"one": [2]int{1, 11},
		"two": [...]int{2, 22},
	}
	fmt.Println("myMap2 = ", myMap2)

	// map of slices
	myMap3 := map[string][]float64{
		"one": []float64{1.1, 1.2, 1.3},
		"two": []float64{2.0},
	}
	fmt.Println("myMap3 = ", myMap3)
}