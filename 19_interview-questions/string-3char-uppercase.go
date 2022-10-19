package main

import (
	"fmt"
	"unicode"
)

func CapitalizeEvery3AlNumChar(s string) string {
	modifiedString := []rune(s)
	count := 0
	for index, eachRune := range modifiedString {
		if unicode.IsNumber(eachRune) || unicode.IsLetter(eachRune) {
			count++
			if count == 3 {
				modifiedString[index] = unicode.ToTitle(eachRune)
				count = 0
			} else {
				modifiedString[index] = unicode.ToLower(eachRune)
			}
		}
	}
	return string(modifiedString)
}

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

type input struct {
	givenStr       string
	skipCharsIndex int
	count          int
	resultstr      string
}

func (i *input) GetValueAsRuneSlice() []rune {
	return []rune(i.givenStr)
}

func (i *input) TransformRune(pos int) {
	// change org string in that position
	eachRune := rune(i.givenStr[pos])
	if unicode.IsNumber(eachRune) || unicode.IsLetter(eachRune) {
		i.count++
		if i.count == i.skipCharsIndex {
			i.resultstr += fmt.Sprintf("%s", string(unicode.ToTitle(eachRune)))
			i.count = 0
		} else {
			i.resultstr += fmt.Sprintf("%s", string(unicode.ToLower(eachRune)))
		}
	}
}

func NewSkipString(skipCharIndex int, s string) input {
	return input{givenStr: s, skipCharsIndex: skipCharIndex}
}

func main() {
	fmt.Println("Method 1")
	for _, test := range []struct {
		inputStr       string
		expectedOutput string
	}{
		{inputStr: "Aspiration.com", expectedOutput: "asPirAtiOn.cOm"},
		{inputStr: "Aspirati1n.com", expectedOutput: "asPirAti1n.cOm"},
		{inputStr: "Aspirati1n.com", expectedOutput: "asPirAti N.coM"},
	} {
		fmt.Printf("CapitalizeEvery3AlNumChar(%s) = %s - Expcted:%s\n",
			test.inputStr, CapitalizeEvery3AlNumChar(test.inputStr), test.expectedOutput)
	}

	fmt.Println("\nMethod 2")
	s := NewSkipString(3, "Aspiration.com")
	MapString(&s)
	fmt.Println(s)
}
