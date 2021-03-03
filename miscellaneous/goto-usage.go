package main

import (
	"fmt"
)

func gotoUseCase() {
	for i := 0; i < 30; i++ {
		if i == 2 {
			goto myNum
		}
	}
myNum:
	fmt.Println(87)
}

// goto "label" & call should be in same function.

func main() {
	gotoUseCase()

}
