package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func removeDelimString(str string) string {
	// alphanumeric (== [0-9A-Za-z])
	// \s is a white space character
	regExp := regexp.MustCompile(`[^[:alnum:]\s]`)
	return regExp.ReplaceAllString(str, "")
}

func sortWordsFixed(input string) string {

	var sorted sort.StringSlice
	originals := map[string]string{}
	var final []string

	unsorted := strings.Split(removeDelimString(input), " ")

	// test each word with first uppercase letter
	for k, v := range unsorted {
		if unicode.IsUpper(rune(v[0])) {
			// tag the words with first uppercase with their index number
			// this is to prevent wrong mapping later
			originals[strings.ToLower(v)+"["+strconv.Itoa(k)+"]"] = v
			sorted = append(sorted, strings.ToLower(v)+"["+strconv.Itoa(k)+"]")
		} else {
			sorted = append(sorted, v)
		}
	}

	// sort our words -- but all words start first letter in lowercase
	sort.Strings(sorted)

	// now, let's replaced back the original words with first uppercase character
	for _, v := range sorted {
		if strings.HasSuffix(v, "]") {
			final = append(final, originals[v])
		} else {
			final = append(final, v)
		}
	}

	return strings.Join(final, " ")
}

func sortWords(input string) string {
	var sorted sort.StringSlice

	// clean
	sorted = strings.Split(removeDelimString(input), " ") // convert to slice

	// bug?
	// words with initial uppercase letter will not be sorted!!
	sort.Strings(sorted)

	return strings.Join(sorted, " ")
}

func main() {
	text := `HEre is an Example of how to convert lines of text into a slice in golang, sort them in descending order and then display the sorted output`

	fmt.Println("Words sort with words with first uppercase letter. Notice that the words are not sorted:")
	fmt.Println("Before: ", text)
	fmt.Println("After: ", sortWords(text))

	fmt.Println("============================================")
	fmt.Println("With sortWordsFixed() function, the words with first uppercase letters are now sorted:")
	fmt.Println("Fixed: ", sortWordsFixed(text))
}
