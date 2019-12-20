package main

/*
Purpose: Variadic Functions
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
}
