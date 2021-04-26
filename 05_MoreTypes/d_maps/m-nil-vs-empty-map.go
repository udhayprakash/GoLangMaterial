package main

import (
	"fmt"
)

func main() {
	var nilMap map[string]bool
	fmt.Println("nilMap  =", nilMap)

	emptyMap := map[string]bool{}
	fmt.Println("emptyMap=", emptyMap)

	fmt.Println("nilMap == nil:", nilMap == nil)

	emptyMap["one"] = true
	fmt.Println("\nemptyMap=", emptyMap)

	// nilMap["one"] = true   // panic: assignment to entry in nil map
	fmt.Println("nilMap  =", nilMap)
}
