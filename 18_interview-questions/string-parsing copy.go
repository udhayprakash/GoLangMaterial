package main

import (
	"fmt"
	"regexp"
)

func main() {

	r := regexp.MustCompile(`([a-zA-Z])[%:]{1}(\w+)`)
	matches := r.FindAllStringSubmatch("A%1,B%2,C%13,d:4,e:5,F:2", -1)
	for _, v := range matches {
		fmt.Println(v)
	}
}
