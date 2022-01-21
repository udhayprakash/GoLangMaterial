package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// type mytype int32 // type aliasing

type point struct {
	x, y int
}

func main() {
	// %d - decimal notation
	fmt.Printf("%d\n", 123) // 123

	// %b - binary notation
	fmt.Printf("%b\n", 123) // 1111011

	fmt.Printf("%x\n", 456)        // 1c8
	fmt.Printf("%x\n", "hex this") // 6865782074686973

	fmt.Printf("%f\n", 123.45) // 123.450000
	fmt.Printf("%g\n", 123.45) // 123.45

	fmt.Printf("%e\n", 123.45) // 1.234500e+02
	fmt.Printf("%E\n", 123.45) // 1.234500E+02

	fmt.Printf("%c\n", 33)           // !
	fmt.Printf("%q\n", "\"string\"") // "\"string\""

	fmt.Printf("%s\n", "\"string\"") // "string"
	fmt.Println()

	fmt.Printf("|%6d|%6d|\n", 12, 345)           // |    12|   345|
	fmt.Printf("|%6d|%6d|\n", 12.3, 34.5)        // |%!d(float64=  12.3)|%!d(float64=  34.5)|
	fmt.Printf("|%6f|%6f|\n", 12.3, 34.5)        // 12.300000|34.500000|
	fmt.Printf("|%6.2f|%6.3f|\n", 12.3, 4.5)     // | 12.30| 4.500|
	fmt.Printf("|%10.2f|%10.3f|\n", 12.3, 4.5)   // |     12.30|     4.500|
	fmt.Printf("|%-10.2f|%-10.3f|\n", 12.3, 4.5) // |12.30     |4.500     |

	fmt.Println()

	fmt.Printf("|%6s|%6s|\n", "foo", "b")   // |   foo|     b|
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // |foo   |b     |
	fmt.Println()

	s := fmt.Sprintf("A %s boy!!!", "good")
	fmt.Println(s, strings.ToUpper(s))

	fmt.Fprintf(os.Stderr, "an %s\n", "error") // an error
	fmt.Println()

	p := point{1, 2}
	fmt.Println(p)

	fmt.Printf("%v\n", p)  // {1 2}
	fmt.Printf("%+v\n", p) // {x:1 y:2}
	fmt.Printf("%#v\n", p) // main.point{x:1, y:2}
	fmt.Println()

	fmt.Println("reflect.TypeOf(p):", reflect.TypeOf(p)) //  main.point
	fmt.Printf("%T\n", p)                                // main.point
	fmt.Printf("%t\n", true)                             // true
}
