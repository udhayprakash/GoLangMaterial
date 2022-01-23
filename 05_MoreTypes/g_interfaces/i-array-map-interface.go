package main

import (
	"fmt"
)

var strSlice = []string{"a", "b", "c"}
var strMap = map[string]interface{}{}

var intSlice = []int{1, 2, 3}
var intMap = map[int]string{}

func main() {

	for i := 0; i != 3; i++ {

		fmt.Println(intSlice[i], "\t", strSlice[i])

		intMap[i] = strSlice[i]
		strMap[strSlice[i]] = intMap
	}

	fmt.Println("Integer map : ", intMap)
	fmt.Println("String map : ", strMap)
}
