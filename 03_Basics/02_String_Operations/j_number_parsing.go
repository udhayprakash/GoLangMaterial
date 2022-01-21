package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// int to string conversion
	x := 123
	y := fmt.Sprintf("x = %d", x)
	z := string(x) // gets the char correponding to unicode number

	fmt.Println(x, reflect.TypeOf(x)) // 123 int
	fmt.Println(y, reflect.TypeOf(y)) // x = 123 string
	fmt.Println(z, reflect.TypeOf(z)) // { string

	p := strconv.Itoa(x)
	fmt.Println(p, reflect.TypeOf(p)) // 123 string

	// To format number in different base
	// 2 - binary; 8 -octal 10-decimal 16-hexadecimal
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
	s := fmt.Sprintf("x=%b", x)                 // "x=1111011"
	fmt.Println(s)
	fmt.Println()

	// To format string representing an integer,
	x, err := strconv.Atoi("123") // x is an int
	fmt.Println(x, err)           // 123 <nil>

	x1, err := strconv.Atoi("12.3") // x1 is an int
	fmt.Println(x1, err)            // 0 strconv.Atoi: parsing "12.3": invalid syntax

	z1, err1 := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	fmt.Println(z1, err1)                       // 123 <nil>

	// With `ParseFloat`, this `64` tells how many bits of
	// precision to parse.
	fmt.Println(strconv.ParseFloat("1.234", 64))    // 1.234 <nil>
	fmt.Println(strconv.ParseFloat("1.234 23", 64)) // 0 strconv.ParseFloat: parsing "1.234 23": invalid syntax

	value, err := strconv.ParseFloat("1.234", 64)
	fmt.Println("Value = ", value, "\terror=", err)

	value, err = strconv.ParseFloat("1.2 34", 64)
	fmt.Println("Value = ", value, "\terror=", err)

	// For `ParseInt`, the `0` means infer the base from the
	// string. `64` requires that the result fit in 64 bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// // `ParseInt` will recognize hex-formatted numbers.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// // A `ParseUint` is also available.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` is a convenience function for basic base-10
	// `int` parsing.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// // Parse functions return an error on bad input.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
