package main

import "fmt"

// 153 = 1*1*1 + 5*5*5 + 3*3*3 // 153 is an Armstrong number
// 1	0, 1, 2, 3, 4, 5, 6, 7, 8, 9
// 3	153, 370, 371, 407
// 4	1634, 8208, 9474
// 5	54748, 92727, 93084
// 6	548834
// 7	1741725, 4210818, 9800817, 9926315
func main() {
	var num, tempNum, rightMost, cubicSum int
	fmt.Println("Enter a 3 digit number:")
	fmt.Scanln(&num)

	tempNum = num

	// get the right most digit
	for tempNum != 0 {
		rightMost = tempNum % 10

		cubicSum += rightMost * rightMost * rightMost

		// update the input digit minus the processed rightMost
		tempNum /= 10 // tempNum = tempNum /10
	}
	fmt.Println("cubicSum=", cubicSum)

	if num == cubicSum {
		fmt.Println(num, "is an AMSTRONG number")
	} else {
		fmt.Println(num, "is NOT a AMSTRONG number")
	}

}
