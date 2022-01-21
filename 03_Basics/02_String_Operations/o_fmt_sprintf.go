package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, dept = "Ravi", "Finance"

	s := fmt.Sprintf("%s is working in %s department\n", name, dept)

	fmt.Println(s)
	io.WriteString(os.Stdout, s)

	// another example

	const num1, num2, num3 = 5, 10, 15
	s = fmt.Sprintf("%d + %d = %d\n", num1, num2, num3)

	io.WriteString(os.Stdout, s)
	fmt.Println(s)
}
