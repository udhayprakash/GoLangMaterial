package main

import (
	"fmt"
)

func main() {

	hits := map[string]map[string]string{
		"#$h-ma": map[string]string{
			"AggregatorA0": "counterA0",
			"AggregatorA1": "counterA1",
			"AggregatorA2": "counterA2",
			"AggregatorA3": "counterA3",
		},
		"#$h-mb": map[string]string{
			"AggregatorB0":  "counterB0",
			"AggregatorB1n": "counterB1",
		},
		"#$h-mc": map[string]string{
			"AggregatorC0": "counterC0",
		},
	}

	// accessing the elements in map within map
	fmt.Println(hits["#$h-mb"]["AggregatorBn"])  // empty
	fmt.Println(hits["#$h-mb"]["AggregatorB1n"]) // counterB1

	fmt.Println(hits["#$h-ma"])                 // show all
	fmt.Println(hits["#$h-ma"]["AggregatorA1"]) // counterA1

	// show all
	fmt.Println(hits)
}
