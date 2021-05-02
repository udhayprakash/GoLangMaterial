package main

import (
	"fmt"
	"strings"
)

func main() {

	builder := strings.Builder{}

	builder.WriteString("There")
	builder.WriteString(" are")
	builder.WriteString(" three")
	builder.WriteString(" hawks")
	builder.WriteString(" in the sky")

	fmt.Println(builder.String())
}
