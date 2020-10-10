package main

import "fmt"

func sum(nums ...float64) (result float64){
	for i:=0; i < len(nums); i++ {
		//fmt.Println("i=", i, nums[i])
		result += nums[i]
	}
	fmt.Println("\nnums :", nums)
	fmt.Println("result:", result)
	return result
}

func sum2(nums ...float64) (result float64){
	for _, value := range nums{
		result += value
	}
	fmt.Println("\nnums :", nums)
	fmt.Println("result:", result)
	return
}

func main(){
	sum()
	sum(10)
	sum(10, 20)
	sum(10, 20, 30.0)
	sum(10, 20, 30.0, 234, 32, 32, 324, 21, 213, 213)
	//sum(10, 13,"23") // cannot use "23" (type untyped string) as type float64 in argument to sum

	sum2()
	sum2(10)
	sum2(10, 20)
	sum2(10, 20, 30.0)
	sum2(10, 20, 30.0, 234, 32, 32, 324, 21, 213, 213)
}

// Assignment: Try to get the average of given values