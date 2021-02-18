package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("decompressor('a3b2c4d10a'):", decompressor("a3b2c4d10a"))
}

func decompressor(inputStr string) string {
	outputStr := ""
	digits := ""
	isNum := false

	for _, eachChr := range inputStr {
		isNum, _ = regexp.MatchString(`\d`, string(eachChr))
		if isNum == true {
			digits += string(eachChr)
		} else {
			if digits != "" {
				digitsNum, _ := strconv.Atoi(digits)
				outputStr = string(outputStr[:len(outputStr)-1]) + strings.Repeat(string(outputStr[len(outputStr)-1]), digitsNum)
			}
			outputStr += string(eachChr)
			digits = ""
		}
	}
	return outputStr
}
