package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var A = 123456
	// for A > 0 {
	// 	fmt.Println(A/1e1, A/1e2, A/1e3, A/1e4, A/1e5, A/1e6)
	// 	A = A/10
	// }

	number := 1234567
	slice := strconv.Itoa(number)

	limit := len(slice)/2
	if len(slice) % 2 != 0{
		limit += 1
	} 
	fmt.Println("limit=", limit)

	result := ""
	for i:=0; i < limit; i++{
		result += slice[i:i+1] + slice[len(slice)-i-1:len(slice)-i]
	}
	fmt.Println("result=", result)
	result1, _ := strconv.Atoi(result)
	fmt.Println("result1=", result1)
	


	fmt.Println(3/2, 3 % 2)
	fmt.Println(3/2, 3 % 2)
}

// reverseDigitsInNumber(54321)
// func reverseDigitsInNumber(N int) {
//     var enable_print int;
//     enable_print = N % 10;
//     for N > 0 {
//         if enable_print == 0 && N % 10 != 0 {
//             enable_print = 1;
//         } else if enable_print == 1 {
//             fmt.Print(N % 10);
//         }
//         N = N / 10;
//     }
// }


// fmt.Println("Solution([]int{1, 3, 6, 4, 1, 2}):", Solution([]int{1, 3, 6, 4, 1, 2}))

// func Solution(A []int) int {
// 	sort.Ints(A)
// 	fmt.Println(A)
// 	var previousVal = 1
// 	for currVal := range A{
// 		if currVal > 0 {
// 			if (previousVal - currVal) == 1 {
// 				previousVal = currVal
// 			} else {
// 				return previousVal + 1
// 			}
// 		} 
// 	}
// 	return previousVal 
// }