package main

import "fmt"

/*
Goto statement is highly discouraged in any programming language
because it becomes diffucult to trace the control flow of program*zc
It makes difficult to understand and modify the code.
*/
func main() {
	for i := 0; i < 9; i++ {
		if i == 5 {
			goto end
		}
		fmt.Println("i=", i)
	}
end:
	fmt.Println("You reach goto label `end` here")
}
