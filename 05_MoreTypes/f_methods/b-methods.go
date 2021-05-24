package main

import (
	"fmt"
	"strings"
)

type Name struct {
	First  string
	Middle string
	Last   string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %c. %s", n.First, n.Middle[0], strings.ToUpper(n.Last))
}

func (n Name) Length() int {
	return len(n.First + n.Middle + n.Last)
}
func main() {
	n1 := Name{"Udhay", "Prakash", "Pethakamsetty"}
	fmt.Printf("%s\n", n1.FullName())
	fmt.Printf("%v\n", n1.Length())
}
