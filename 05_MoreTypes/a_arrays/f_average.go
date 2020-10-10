package main

import "fmt"

func main() {
	marks := [6]float32{89, 56, 89, 99, 78, 90}
	fmt.Println("marks    :", marks)
	fmt.Println("summation:", summation(marks))
	fmt.Println("Avaerge  :", average(marks))
}

func summation(nums [6]float32) float32 {
	var result float32

	for _, value := range nums {
		result += value
		fmt.Println("\tvalue:", value)
	}
	return result
}

func average(nums [6]float32) float32 {
	var result float32

	for index := 0; index < len(nums); index++ {
		result += nums[index]
		fmt.Println("\tnums[index]= ", nums[index])

	}

	return result / float32(len(nums))
}
