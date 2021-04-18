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

func removeSpaces(givenStr string) string {
	filteredStr := ""
	tempChar := ""
	for i := 0; i < len(givenStr); i++ {
		tempChar = string(givenStr[i])
		if tempChar != string(" ") {
			filteredStr += tempChar
		}
	}
	return filteredStr
}

func main() {
	fmt.Println(`isPalindrome("hello"):`, isPalindrome("hello"))
	fmt.Println(`isPalindrome("ollo") :`, isPalindrome("ollo"))
	fmt.Println(`isPalindrome("racecar") :`, isPalindrome("racecar"))

	fmt.Println(`isPalindrome("a man a plan a canal panama") :`, isPalindrome("a man a plan a canal panama"))
	fmt.Println(`isPalindrome(removeSpaces("a man a plan a canal panama")) :`,
		isPalindrome(removeSpaces("a man a plan a canal panama")))
}
