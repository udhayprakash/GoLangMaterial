package palindrome

import "fmt"

func isPalindrome(givenString string) bool {
	var reverseString string
	strLen := len(givenString)

	for i := 0; i < strLen; i++ {
		reverseString += string(givenString[strLen-i-1])
	}
	return reverseString == givenString
}

func main() {
	fmt.Println(`isPalindrome("hello"):`, isPalindrome("hello"))
	fmt.Println(`isPalindrome("ollo") :`, isPalindrome("ollo"))
}
