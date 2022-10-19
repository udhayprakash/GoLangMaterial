package main

import (
	"fmt"
	"strings"
)

func checkBalanceBraces(expression string) bool {
	braces := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	bracesFound := []rune{}
	for _, eachChr := range expression {
		if strings.Contains("{([", string(eachChr)) {
			bracesFound = append(bracesFound, eachChr)
		} else if strings.Contains("})]", string(eachChr)) {
			if len(bracesFound) == 0 || bracesFound[len(bracesFound)-1] != braces[eachChr] {
				return false
			}
			bracesound = bracesFound[:len(bracesFound)-1]
		}
	}
	if len(bracesFound) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(checkBalanceBraces("(())") == true)
	fmt.Println(checkBalanceBraces("[](){") == false)
	fmt.Println(checkBalanceBraces("sdjf{[df]}()") == true)
	fmt.Println(checkBalanceBraces("sdjf{[d)f]})()") == false)
	fmt.Println(checkBalanceBraces("sdjf{[d()f]})()") == false)
	fmt.Println(checkBalanceBraces("sdjf{[d()f]}()") == true)
	fmt.Println(checkBalanceBraces("{{") == false)
	fmt.Println(checkBalanceBraces("{{()") == false)
	fmt.Println(checkBalanceBraces("{[()}]") == false)
	fmt.Println(checkBalanceBraces("[{()()}({[]})]({}[({})])((((((()[])){}))[]{{{({({({{{{{{}}}}}})})})}}}))[][][]{") == false)
	fmt.Println(checkBalanceBraces("[{()()}({[]})]({}[({})])((((((()[])){}))[]{{{({({({{{{{{}}}}}})})})}}}))[][][]]") == false)
	fmt.Println(checkBalanceBraces("[{()([)}({[]})]({}[({})])((((((()[])){}))[]{{{({({({{{{{{}}}}}})})})}}}))[][][]") == false)
	fmt.Println(checkBalanceBraces("{}[]()") == true)
	fmt.Println(checkBalanceBraces("{[}]") == false)
	fmt.Println(checkBalanceBraces(")") == false)
	fmt.Println(checkBalanceBraces("}") == false)
	fmt.Println(checkBalanceBraces("]") == false)
	fmt.Println(checkBalanceBraces("[()]") == true)
	fmt.Println(checkBalanceBraces("[{}]") == true)
	fmt.Println(checkBalanceBraces("{][}") == false)
}
