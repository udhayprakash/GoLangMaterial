package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// Int to String conversion
	x := 97

	// Method 1
	z := string(x)
	fmt.Println(x, reflect.TypeOf(x)) // 97 int
	fmt.Println(z, reflect.TypeOf(z)) // a string

	// Method 2  - digit as it is in string
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, reflect.TypeOf(y)) // 97 string

	// Method 3  - digit as it is in string
	p := strconv.Itoa(x)
	fmt.Println(p, reflect.TypeOf(p)) // 97 string
	fmt.Println()

	// To format number in different base
	// 2 - binary; 8 -octal 10-decimal 16-hexadecimal
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1100001"
	fmt.Println(fmt.Sprintf("%b", x))           // "1100001"

	fmt.Println(strconv.FormatInt(int64(x), 8)) // "0o141"
	fmt.Println(fmt.Sprintf("%o", x))           // "0o141"

	fmt.Println(strconv.FormatInt(int64(x), 10)) // "97"
	fmt.Println(fmt.Sprintf("%d", x))            // "97"

	fmt.Println(strconv.FormatInt(int64(x), 16)) // "0x61"
	fmt.Println(fmt.Sprintf("%X", x))            // "0x61"
	fmt.Println()

	// ==== To format string representing an integer,
	fmt.Println(strconv.Atoi("97")) // 97 <nil>

	x, err := strconv.Atoi("123") // x is an int
	fmt.Println(x, err)           // 123 <nil>

	x1, err := strconv.Atoi("12.3") // x1 is an int
	fmt.Println(x1, err)            // 0 strconv.Atoi: parsing "12.3": invalid syntax

	x, err = strconv.Atoi("123Helo") // x is an int
	fmt.Println(x, err)              // 0 strconv.Atoi: parsing "123

	// `Atoi` is a convenience function for basic base-10
	// `int` parsing.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse functions return an error on bad input.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)

	// For `ParseInt`, the `0` means infer the base from the
	// string. `64` requires that the result fit in 64 bits.
	z1, err1 := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	fmt.Println(z1, err1)                       // 123 <nil>

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` will recognize hex-formatted numbers.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// With `ParseFloat`, this `64` tells how many bits of
	// precision to parse.
	fmt.Println(strconv.ParseFloat("1.234", 64))    // 1.234 <nil>
	fmt.Println(strconv.ParseFloat("1.234 23", 64)) // 0 strconv.ParseFloat: parsing "1.234 23": invalid syntax

	value, err := strconv.ParseFloat("1.234", 64)
	fmt.Println("Value = ", value, "\terror=", err)

	value, err = strconv.ParseFloat("1.2 34", 64)
	fmt.Println("Value = ", value, "\terror=", err)

	// A `ParseUint` is also available.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
}
