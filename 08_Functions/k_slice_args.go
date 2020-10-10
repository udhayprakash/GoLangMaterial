package main

import "fmt"

func find(num int, nums []int) {
	fmt.Printf("\n\ntype of nums is %T\n", nums)
	found := false
	for index, value := range nums {
		if value == num {
			fmt.Printf("%d found at index %d\n", value, index)
			found = true
		}
	}
	if found == false {
		fmt.Printf("%d not found in %d\n", num, nums)
	}
}

func main() {
	find(89, []int{89, 90, 95})
	find(45, []int{56, 67, 45, 90, 109})
	find(78, []int{38, 56, 98})
	find(87, []int{})
}
