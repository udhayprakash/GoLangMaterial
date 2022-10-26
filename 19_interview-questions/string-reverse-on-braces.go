package main

import (
	"fmt"
	"sort"
)

/*
Purpose: Reverse the string between the braces.

	First reverse string within the inner-most braces, followed by outermost

	ASSUMPTION - all braces WERE balanced, in given string
*/
func reverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	/*
		ABC((xyz)def) ==> ABCfedxyz
			ABC(zyxdef)
		     ABCfedxyz

		map[3:12 4:8]

			A	B	C	(	(	x	y	z	)	d	 e	  f		)
			0   1   2   3   4   5   6   7   8   9   10   11		12

			A	B	C	(		z	y	x		d	 e	  f		)
			0   1   2   3       4   5   6       7    8    9		10
	*/
	fmt.Println(ReverseSubStringWithInBraces("ABC((xyz)def)") == "ABCfedxyz")
	fmt.Println(ReverseSubStringWithInBraces("ABC(xyz)def") == "ABCzyxdef")
	fmt.Println(ReverseSubStringWithInBraces("ABC(xyz)(def)") == "ABCzyxfed")
	fmt.Println(ReverseSubStringWithInBraces("(ABC)(xyz)(def)") == "CBAzyxfed")

}

func ReverseSubStringWithInBraces(givenStr string) string {
	var bracePos []int
	bracePair := make(map[int]int)
	var startIndices []int
	for index, char := range givenStr {
		if string(char) == "(" {
			bracePos = append(bracePos, index)
			startIndices = append(startIndices, index)
		} else if string(char) == ")" {
			lastPos := len(bracePos) - 1
			bracePair[bracePos[lastPos]] = index
			bracePos = bracePos[:lastPos]
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(startIndices)))
	prevStartIndex := 0
	for index, startIndex := range startIndices {
		endIndex, _ := bracePair[startIndex]
		if index != 0 && endIndex > prevStartIndex {
			// Needed for nexted braces condition
			endIndex -= index * 2
		}
		prevStartIndex = startIndex
		givenStr = givenStr[:startIndex] + reverseStr(givenStr[startIndex+1:endIndex]) + givenStr[endIndex+1:]
	}
	return givenStr
}
