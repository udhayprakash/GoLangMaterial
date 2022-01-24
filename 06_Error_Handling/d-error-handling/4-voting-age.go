package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age:")
	fmt.Scanf("%d", &age)
	voting(age)
	fmt.Println("Execution flow completed without any error")
}

func voting(voterAge int) {
	defer ageRestriction()
	if voterAge < 18 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", voterAge))
	}
	fmt.Println("You are Eligible for voting")
}

func ageRestriction() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in f", r)
	}
	fmt.Println("Minimum age for voting is 18!!!")
}
