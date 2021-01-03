package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	str1 := "123"
	checkNumeric(str1)
	fmt.Println(str1, IsNumeric(str1))

	str2 := "12a3"
	checkNumeric(str2)
	fmt.Println(str2, IsNumeric(str2))

	str3 := "abcd"
	checkNumeric(str3)
	fmt.Println(str3, IsNumeric(str3))

	str4 := "12 3" // space will cause the "standard way" test to fail, trim space first if needed

	checkNumeric(str4)                 // more robust, this is what I need
	fmt.Println(str4, IsNumeric(str4)) // the "standard way" will mark this is false

	str41 := "12 3 45 6 7"

	checkNumeric(str41)
	fmt.Println(str41, IsNumeric(str41))

	str43 := "12 3 a 45 6 7"

	checkNumeric(str43)
	fmt.Println(str43, IsNumeric(str43))

	str5 := "12.3"
	checkNumeric(str5)
	fmt.Println(str5, IsNumeric(str5))

	str6 := "-123.4"
	checkNumeric(str6)
	fmt.Println(str6, IsNumeric(str6))
}

func checkNumeric(input string) {
	var def bool = false

	for _, v := range input {
		if unicode.IsDigit(v) {
			def = true
		} else if unicode.IsPunct(v) {
			continue // skip
		} else if unicode.IsSpace(v) {
			continue // skip
		} else {
			def = false
			break
		}
	}

	fmt.Println(input, def)
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
