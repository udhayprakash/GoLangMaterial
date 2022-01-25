package main

import "fmt"

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func getAllOddNumbers(nums []int) []int {
	var resultNums []int
	for _, num := range nums {
		if isOdd(num) == true {                // function
			resultNums = append(resultNums, num)
		}
	}
	return resultNums
}

func getAllEvenNumbers(nums []int) []int {
	var resultNums []int
	for _, num := range nums {
		if isEven(num) == true {                // function
			resultNums = append(resultNums, num)
		}
	}
	return resultNums
}

// Decorator design pattern - passing func as argument 
type EvenOddFunc func(int) bool

func Filter(nums []int, myfunc EvenOddFunc) []int {
	var resultNums []int
	for _, num := range nums {
		if myfunc(num) == true {                // function
			resultNums = append(resultNums, num)
		}
	}
	return resultNums
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)

	fmt.Println("isOdd(2) =", isOdd(2))
	fmt.Println("isOdd(3) =", isOdd(3))
	fmt.Println()
	
	fmt.Println("getAllOddNumbers(slice)  =", getAllOddNumbers(slice))
	fmt.Println("getAllEvenNumbers(slice) =", getAllEvenNumbers(slice))
	fmt.Println()

	fmt.Println("Filter(slice, isEven)    =", Filter(slice, isEven))
	fmt.Println("Filter(slice, isOdd)     =", Filter(slice, isOdd))
}
