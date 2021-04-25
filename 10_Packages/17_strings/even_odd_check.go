package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter a number to test if odd or even : ")
	//fmt.Scanf("%d", &num)

	numberReader := bufio.NewReader(os.Stdin)
	num, _ := numberReader.ReadString('\n')

	// the input can be multiple integers with spaces in between
	// turn to slice

	numSlice := strings.Fields(num)

	for _, v := range numSlice {
		i, _ := strconv.Atoi(v)
		if i%2 == 0 {
			fmt.Println(v, "is Even")
		} else {
			fmt.Println(v, "is Odd")
		}
	}
	
	
	// Fields function also splits the string
    // on the basis of white spaces and tabs
    s := strings.Fields(" I \t love \n to \n code \n all \t day.")
    fmt.Println(s)
  
    // Splits into 5 words which have
    // newline character in between
    s = strings.Fields("I\nam\nlearning\nGo\nlanguage")
    fmt.Println(s)
}
