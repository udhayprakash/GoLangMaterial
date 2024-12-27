package main

import (
	"fmt"
	"reflect"
)

// comparable types: boolean, numeric, string, pointer,
//                    channel, and interface types

func main() {
	choice1, choice2 := true, false
	fmt.Println("choice1 == choice2 :", choice1 == choice2)

	num1, num2 := 12, 34
	fmt.Println("num1 == num2       :", num1 == num2)

	name1, name2 := "Ravi", "ravi"
	fmt.Println("name1 == name2     :", name1 == name2)

	// maps are not comparable objects
	map1 := map[string]int{}
	map2 := map1
	// fmt.Println("map1 == map2     :", map1 == map2)
	// invalid operation: map1 == map2 (map can only be compared to nil)

	fmt.Println("map1 == nil     :", map1 == nil)
	fmt.Println("map2 == nil     :", map2 == nil)

	fmt.Println("reflect.DeepEqual(map1, map2):", reflect.DeepEqual(map1, map2))

	map3 := make(map[string]int)
	fmt.Println("reflect.DeepEqual(map1, map3):", reflect.DeepEqual(map1, map3))
}


// Assignment - compare the map objecs, of multidimensionals
