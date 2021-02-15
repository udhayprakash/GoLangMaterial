// Golang program to illustrate the usage of
// fmt.Sprintf() function

// Including the main package
package main

// Importing fmt, io and os
import (
	"fmt"
	"io"
	"os"
)

// Calling main
func main() {

	// Declaring some const variables
	const name, dept = "GeeksforGeeks", "CS"

	// Calling Sprintf() function
	s := fmt.Sprintf("%s is a %s Portal.\n", name, dept)

	// Calling WriteString() function to write the
	// contents of the string "s" to "os.Stdout"
	io.WriteString(os.Stdout, s)
	fmt.Println(s)

	// Declaring some const variables
	const num1, num2, num3 = 5, 10, 15

	// Calling Sprintf() function
	s = fmt.Sprintf("%d + %d = %d\n", num1, num2, num3)

	// Calling WriteString() function to write the
	// contents of the string "s" to "os.Stdout"
	io.WriteString(os.Stdout, s)
	fmt.Println(s)
}
