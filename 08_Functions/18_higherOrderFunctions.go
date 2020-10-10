package main

import (
	"fmt"
	"strings"
)
// strings.Map applies a function to each character of a string,
// joining the results to make another string
func add2(r rune) rune { return r + 1 }

func ceaserCipher(r rune) rune { return r + 3 }

func main() {

	fmt.Println(strings.Map(add2, "HAL-9000"))      // "IBM.:111"
	fmt.Println(strings.Map(add2, "VMS"))           // "WNT"
	fmt.Println(strings.Map(add2, "Admix"))         // "Benjy"

	fmt.Println(strings.Map(ceaserCipher, "Admix")) // "Dgpl{"
}
