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

func main() {
	n1 := Name{"Udhay", "Prakash", "Pethakamsetty"}
	fmt.Printf("%s", n1.FullName())
}

// NOTE: You can only define methods on types within the same package.
