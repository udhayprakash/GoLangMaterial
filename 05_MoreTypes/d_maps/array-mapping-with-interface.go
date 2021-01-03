package main

import (
	"fmt"
)

var strArray = []string{"abc", "def", "ghi"}
var strMap = map[string]interface{}{}

var intArray = []int{1, 2, 3}
var intMap = map[int]string{}

func main() {

	for i := 0; i != 3; i++ {

		fmt.Println(intArray[i], "\t", strArray[i])

		intMap[i] = strArray[i]
		strMap[strArray[i]] = intMap
	}

	fmt.Println("String map : ", strMap)
	fmt.Println("Integer map : ", intMap)
}
