package main

import "fmt"

func printRow(ch byte, numcols int) {
	for i := 0; i < numcols; i++ {
		fmt.Printf("%c ", ch)
	}
}

func printPattern(n int) {
	var ch byte = '*'
	for i := 0; i < n; i++ {
		printRow(ch, n)
		fmt.Println()
		if ch == '*' {
			ch = '+'
		} else {
			ch = '*'
		}
	}
}

func main() {
	n := 5
	printPattern(n)
}