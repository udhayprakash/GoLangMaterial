package main

import (
	"fmt"
)

func CountDigits(i int) (count int) {
	for i != 0 {

		i /= 10
		count = count + 1
	}
	return count
}

func main() {

	var i int
	fmt.Println("Enter an integer value : ")

	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("You have entered a : ", CountDigits(i), "digit(s) integer value")

}
