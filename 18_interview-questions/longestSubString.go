package main

import (
	"fmt"
)

func lengthOfLongestSubstring(str string) int {
	str_len := len(str)
	arr := make([]bool, 125)
	max := 1
	count := 0
	// If string length is zero
	if str_len == 0 {
		max = 0
	} else {
		for i := 0; i < str_len; i++ {
			for j := 0; j < 125; j++ {
				arr[j] = false
			}
			count = 0
			arr[str[i]] = true
			count = 1

			for k := i + 1; k < str_len; k++ {
				if arr[str[k]] {
					break
				}

				arr[str[k]] = true
				count = count + 1

				if max < count {
					max = count
				}
			}
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
}
