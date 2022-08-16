package anagram2

import (
	"reflect"
)

func GetCharCount(s string) (c map[rune]int) {
	c = make(map[rune]int)
	for _, runeValue := range s {
		if _, ok := c[runeValue]; ok {
			c[runeValue] += 1
		} else {
			c[runeValue] = 1
		}
	}
	return
}

func AreAnagram(s1, s2 string) bool {
	c1 := GetCharCount(s1)
	c2 := GetCharCount(s2)

	return reflect.DeepEqual(c1, c2)
}
