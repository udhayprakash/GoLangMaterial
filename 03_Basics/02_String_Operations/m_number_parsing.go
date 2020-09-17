package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"

	// To format number in different base
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"

	s := fmt.Sprintf("x=%b", x) // "x=1111011"
	fmt.Println(s)

	//// To format string representing an integer,
	//x, err := strconv.Atoi("123") // x is an int
	//fmt.Println(x, err)
	//y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	//fmt.Println(y, err)
}
