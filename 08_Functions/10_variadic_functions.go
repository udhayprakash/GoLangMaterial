package main

/*
Purpose: Variadic Functions
	- http://en.wikipedia.org/wiki/Variadic_function
	- Function with variable arguments
	- can be called with any number of trailing arguments.
	ex: fmt.Println(), append()

In the declaration of the variadic function,
the type of the last parameter is preceded by
an ellipsis, i.e, (…). It indicates that the
function can be called at any number of
parameters of this type.

When to use
------------
	- when you want to pass a slice in a function.
	- when we don’t know the quantity of parameters.
It increase the readability of your program.
*/

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum()
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	//sum(nums) // cannot use nums (type []int) as type int in argument to sum
	sum(nums...)

	fmt.Printf("%T\n", f1) // "func(...int)"
	fmt.Printf("%T\n", f2) // "func([]int)"
}

func f1(...int) {}
func f2([]int)  {}


// Assignments
// 1. variadic functions max & min, which accepts minimum 1 argument
// 2. variadic function version of strings.Join