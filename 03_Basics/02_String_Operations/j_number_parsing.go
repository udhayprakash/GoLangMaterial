package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"
	fmt.Println()

	// To format number in different base
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
	s := fmt.Sprintf("x=%b", x)                 // "x=1111011"
	fmt.Println(s)
	fmt.Println()

	// To format string representing an integer,
	x, err := strconv.Atoi("123") // x is an int
	fmt.Println(x, err)
	z, err1 := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	fmt.Println(z, err1)

	// With `ParseFloat`, this `64` tells how many bits of
	// precision to parse.
	fmt.Println(strconv.ParseFloat("1.234", 64))
	fmt.Println(strconv.ParseFloat("1.234 23", 64))

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
