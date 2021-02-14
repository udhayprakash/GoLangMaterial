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
	var value string
	for i := 0; i <= 100; i++ {
		value = ""
		if i%3 == 0 {
			value += "Fizz"
		}
		if i%5 == 0 {
			value += "Buzz"
		}
		if value == "" {
			fmt.Println(i)
		} else {
			fmt.Println(value)
		}
	}
}
