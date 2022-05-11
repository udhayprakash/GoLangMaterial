package main

import (
	"fmt"
	"strings"
)

func JoinLines(input []string, joinBy string) string {
	return strings.Join(input, joinBy)
}

func OnlyJoinLines(input []string, onlyJoinBy string) string {

	var result []string

	var skipIndex int

	for k, v := range input {
		if strings.HasSuffix(v, onlyJoinBy) {
			// only join lines that match the criteria

			if k+1 >= len(input) {
				result = append(result, v+input[k]+"\n")
				skipIndex = k
			} else {
				result = append(result, v+input[k+1]+"\n")
				// after joining, mark the next index to skip
				skipIndex = k + 1
			}
		} else {
			if !(skipIndex > 0) {
				result = append(result, v+"\n")
			}
			// remember to reset back to 0
			skipIndex = 0
		}

	}

	return strings.Join(result, "")

}

func main() {
	lines := `line 1
 line 2
 line 3:
 line 4
 line 5:
 line 6
 line 7
 line 8+
 line 9`

	// turn lines into slice
	lineSlice := strings.Split(lines, "\n")

	result1 := JoinLines(lineSlice, " ")
	fmt.Println("RESULT 1 : ", result1)

	// join all lines with :
	result2 := JoinLines(lineSlice, ":")
	fmt.Println("RESULT 2 : ", result2)

	// ONLY join lines that have : symbol as suffix
	result3 := OnlyJoinLines(lineSlice, ":")
	fmt.Println("RESULT 3 : ", result3)

	// ONLY join lines that have + symbol as suffix
	result4 := OnlyJoinLines(lineSlice, "+")
	fmt.Println("RESULT 4 : ", result4)

}
