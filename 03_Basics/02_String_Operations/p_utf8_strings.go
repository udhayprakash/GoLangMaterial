package main

import (
	"fmt"
	"unicode/utf8"
)

func RemoveLastChar(str string) string {
	for len(str) > 0 {
		_, size := utf8.DecodeLastRuneInString(str)
		return str[:len(str)-size]
	}
	return str
}

func main() {
	fmt.Println(`RemoveLastChar("battle") :`, RemoveLastChar("battle"))
	fmt.Println(`RemoveLastChar("battle3"):`, RemoveLastChar("battle3"))
	fmt.Println(`RemoveLastChar("καλημ")  :`, RemoveLastChar("καλημ ρα κóσμ"))
}
