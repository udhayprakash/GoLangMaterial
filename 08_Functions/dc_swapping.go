package main

import "fmt"

func swap(w1, w2 string) (string, string) {
	return w2, w1
}

func main() {
	// word1 := swap("Hello", "world")
	// 1 variable but swap returns 2 values

	word1, word2 := swap("Hello", "world")
	fmt.Println(word1, word2)
}
