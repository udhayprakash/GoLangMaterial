package main

import "fmt"

/*
Problem: Write a program that prints the numbers from
1 to 100. But for multiples of three print, “Fizz”
instead of the number, and for multiples of five,
print “Buzz”. For numbers which are multiples of both
three and five, print “FizzBuzz”.
*/
func main() {
	fizzBuzz(15)
}

func fizzBuzz(n int32) {
	var i int32
	var output string
	for i = 1; i <= n; i++ {
		output = ""
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}
		if output == "" {
			output = fmt.Sprintf("%d", i)
		}
		fmt.Println(output)
	}

}
