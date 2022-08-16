package main

import "fmt"

func stringReverse(strs []string) []string {
	l := len(strs) - 1
	for i := 0; i < l; i++ {
		strs[i], strs[l-i] = strs[l-i], strs[i]
	}
	return strs

}

func main() {
	s1 := []string{"this", "is", "a", "very", "great", "day"}
	fmt.Println(stringReverse(s1)) // [day is a very great this]
}
