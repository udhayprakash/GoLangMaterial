package main

import "fmt"

func myfunc(searchStr, paragraph string) int {
	for i := 0; i < len(paragraph); i++ {
		found := true
		for index, char := range searchStr {
			if string(char) != string(paragraph[i+index]) {
				found = false
				break
			}
		}
		if found == true {
			return i
		}

	}
	return -1
}

func main() {
	sentence := "python is a cool language"

	fmt.Println("myfunc(\"ool\", sentence)      =", myfunc("ool", sentence))      // 13
	fmt.Println("myfunc(\"java\", sentence)     =", myfunc("java", sentence))     // -1
	fmt.Println("myfunc(\"python\", sentence)   =", myfunc("python", sentence))   // 0
	fmt.Println("myfunc(\"language\", sentence) =", myfunc("language", sentence)) // 17

}
