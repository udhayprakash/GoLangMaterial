package main

import (
	"fmt"
	"os"
)

func mapper(placeholder string) string {
	switch placeholder {
	case "FRUIT":
		return "mango"
	case "CAR":
		return "audi"
	default:
		return "<empty>"
	}
}

func main() {

	// raw string
	raw := "I am eating $FRUIT and driving ${CAR}."

	// formatted string
	formatted := os.Expand(raw, mapper)

	// print formatted string
	fmt.Println(formatted)

}
