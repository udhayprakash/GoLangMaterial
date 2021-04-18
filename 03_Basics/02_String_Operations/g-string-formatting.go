package main

import (
	"fmt"
	"os"
	"reflect"
)

//type mytype int32

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Println(p)

	fmt.Printf("%v\n", p)  // {1 2}
	fmt.Printf("%+v\n", p) // {x:1 y:2}
	fmt.Printf("%#v\n", p) // main.point{x:1, y:2}
	fmt.Println()

	fmt.Println("reflect.TypeOf(p):", reflect.TypeOf(p))
	fmt.Printf("%T\n", p)    // main.point
	fmt.Printf("%t\n", true) // true

	// %d - decimal notation
	fmt.Printf("%d\n", 123) // 123

	// %b - binary notation
	fmt.Printf("%b\n", 14) // 1110

	fmt.Printf("%c\n", 33) // !

	fmt.Printf("%x\n", 456) // 1c8

	fmt.Printf("%f\n", 78.9) // 78.900000

	fmt.Printf("%e\n", 123400000.0) // 1.234000e+08
	fmt.Printf("%E\n", 123400000.0) // 1.234000E+08

	fmt.Printf("%s\n", "\"string\"") // "string"

	fmt.Printf("%q\n", "\"string\"") // "\"string\""

	fmt.Printf("%x\n", "hex this") // 6865782074686973

	fmt.Printf("%p\n", &p) // 0xc000064090
	fmt.Println()

	fmt.Printf("|%6d|%6d|\n", 12, 345)           // |    12|   345|
	fmt.Printf("|%6d|%6d|\n", 12.3, 34.5)        // |%!d(float64=  12.3)|%!d(float64=  34.5)|
	fmt.Printf("|%6f|%6f|\n", 12.3, 34.5)        // |12.300000|34.500000|
	fmt.Printf("|%6.2f|%6.3f|\n", 12.3, 4.5)     // | 12.30| 4.500|
	fmt.Printf("|%10.2f|%10.3f|\n", 12.3, 4.5)   // |     12.30|     4.500|
	fmt.Printf("|%-10.2f|%-10.3f|\n", 12.3, 4.5) // |12.30     |4.500     |
	fmt.Println()

	fmt.Printf("|%6s|%6s|\n", "foo", "b")   // |   foo|     b|
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // |foo   |b     |
	fmt.Println()

	s := fmt.Sprintf("A %s boy!!!", "good")
	fmt.Println(s) // A good boy!!!

	fmt.Fprintf(os.Stderr, "an %s\n", "error") // an error
}
