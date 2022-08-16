package main

import (
	"fmt"
	"regexp"
	"strings"
	"text/scanner"
	"unicode"
)

// faster way to get rid delimiters
// for string without unicode

func removeDelimString(str string) string {
	// alphanumeric (== [0-9A-Za-z])
	// \s is a white space character
	regExp := regexp.MustCompile(`[^[:alnum:]\s]`)
	return regExp.ReplaceAllString(str, "")
}

// get all upper case characters from string
// to get all lower case characters, just remove the exclamation mark !
// infront of unicode.IsLower()

func getUpperCaseChars(str string) []string {

	tokens := removeDelimString(str)

	var result []string

	for _, char := range tokens {

		if !unicode.IsLower(char) && char != ' ' {
			fmt.Println(scanner.TokenString(char))
			result = append(result, scanner.TokenString(char))
		}
	}

	return result
}

func main() {
	str := "This a string with some UpperCase Characters."

	temp := getUpperCaseChars(str)

	for k, v := range temp {
		fmt.Println(k, v)
	}

	// form into a UPPERCASE string

	upper := removeDelimString(strings.Join(temp, ""))
	fmt.Println(upper)

}
