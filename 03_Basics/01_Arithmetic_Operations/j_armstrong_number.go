package main

import (
	"fmt"
)

func main() {

	var rightMost, num int
	var cubicSum int = 0
	var tempNum int = 0

	fmt.Print("Enter a 3 digits number : ")
	fmt.Scanf("%d", &num)

	tempNum = num

	// get the right most digit
	for {
		rightMost = tempNum % 10
		cubicSum += rightMost * rightMost * rightMost

		// update the input digit minus the processed rightMost
		tempNum /= 10

		if tempNum == 0 {
			// break the for loop
			break
		}
	}

	if num == cubicSum {
		fmt.Println(num, "is an Armstrong number!")
	} else {
		fmt.Println(num, "is NOT an Armstrong number!")
	}
}
