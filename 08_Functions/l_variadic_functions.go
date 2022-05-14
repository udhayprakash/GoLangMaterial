package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
variadic functions
	- functions which support any no. of arguments
	- spread operator
*/
func NoArgsFunc() {
	fmt.Println("No Arguments")
}

func oneArgFunc(num int) {
	fmt.Println("Args:", num)
}

func TwoArgFunc(num1, num2 int) {
	fmt.Println("Args:", num1, num2)
}

// variadic functions
func AnyNoOfArgs(num ...int) {
	fmt.Println("Args:", num)
}

func AnyNoOfArgsAnyType(num ...interface{}) {
	fmt.Println("Args:", num)
}

func main() {
	NoArgsFunc()
	oneArgFunc(111) // Args: 111
	TwoArgFunc(111, 222)

	// Variadic function Ex: fmt.Println()
	fmt.Println()
	fmt.Println(1)
	fmt.Println(1, 2)
	fmt.Println(1, 2, 3)
	fmt.Println(1, 2, 3, 4, 5, 6, 7, 8)

	fmt.Println(1, 2.3, true, nil, []int{7, 8}, map[int]int{1: 1, 2: 2})

	AnyNoOfArgs()
	AnyNoOfArgs(1)
	AnyNoOfArgs(1, 2)
	AnyNoOfArgs(1, 2, 3, 5, 6, 3, 2, 213)

	AnyNoOfArgsAnyType(1, '2', "udhay", true, [3]int{1, 2, 3})
	fmt.Println()

	SumAvg(67)
	SumAvg(12, 34, 67)
	fmt.Println()

	getFullName("udhay")
	getFullName("udhay", "prakash")
	getFullName("udhay", "prakash", "pethakamsetty")
	fmt.Println()

	fmt.Println(createSentence([]string{"udhay"}...))
	fmt.Println(createSentence([]string{"udhay", "prakash"}...))
	fmt.Println(createSentence([]string{"udhay", "prakash", "pethakamsetty"}...))
}

func SumAvg(nums ...int) {
	summation := 0
	for _, num := range nums {
		summation += num
	}
	fmt.Printf("Sum(%d) = %d\n", nums, summation)
	fmt.Printf("Avg(%d) = %f\n", nums, float32(summation)/float32(len(nums)))
}

func getFullName(name_parts ...string) {
	fullName := ""
	for _, name := range name_parts {
		fullName += name
	}
	fmt.Printf("getFullName(%v)=%v\n", name_parts, fullName)
}

func createSentence(words ...string) string {
	fmt.Println(reflect.TypeOf(words)) // []string
	return strings.Join(words, " ")
}

// Assignment: Try to compute mean, median and mode
