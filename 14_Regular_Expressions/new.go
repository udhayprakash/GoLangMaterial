package main

import (
	"fmt"

	"regexp"
)


func main() {

	re := regexp.MustCompile(`(\d+)\sViews`)

	match := re.FindString("100 Views")

	fmt.Println("match:", match)
}
