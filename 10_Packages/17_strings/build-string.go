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

	fmt.Printf("Capacity:%d\t Length:%d\n", builder.Cap(), builder.Len())

	// Grow the capacity, if we know, in advance, how much data
	// is about to be written into buffer, to minimize copies
	builder.Grow(1000)
	fmt.Printf("Capacity:%d\t Length:%d\n", builder.Cap(), builder.Len())
	fmt.Println()

	// Consuming the build string
	for i := 0; i <= 10; i++ {
		fmt.Fprint(&builder, "\nstring %d ---", i)
	}
	fmt.Println("builder.String():", builder.String())

	// To reset the string builder
	builder.Reset()
	fmt.Println("After reset, ")
	fmt.Printf("Capacity:%d\t Length:%d\n", builder.Cap(), builder.Len())
}
