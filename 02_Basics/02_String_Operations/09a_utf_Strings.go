package main

import (
	"fmt"
	"golang.org/x/text/cases"
)

func main()  {
	c := cases.Fold()
	fmt.Printf("%s %v", c, c.String("grüßen"))
}