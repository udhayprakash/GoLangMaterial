package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	fmt.Printf("")

	// %d - decimal notation
	fmt.Printf("%d\n", 123) // 123

	// %b - binary notation
	fmt.Printf("%b\n", 123) // 1111011

	// Hexadecimal - 0 to 9, A to Z
	fmt.Printf("%x\n", 456)        // 1c8
	fmt.Printf("%x\n", "hex this") // 6865782074686973

	fmt.Printf("%f\n", 123.45) // 123.450000
	fmt.Printf("%g\n", 123.45) // 123.45

	fmt.Printf("%e\n", 123.45) // 1.234500e+02
	fmt.Printf("%E\n", 123.45) // 1.234500E+02

	// rune
	fmt.Printf("%c\n", 33) // !
	fmt.Printf("%q\n", 33) // '!'

	// string
	fmt.Printf("%c\n", '!') // !
	fmt.Printf("%s\n", '!') // %!s(int32=33)
	fmt.Printf("%s\n", "!") // !
	fmt.Println()

	// formating with digit space reserved
	fmt.Printf("|%6d |%6d|\n", 12, 12)     // |    12 |    12|
	fmt.Printf("|%6d |%6d|\n", 12.3, 34.5) // |%!d(float64=  12.3) |%!d(float64=  34.5)|

	fmt.Printf("|%6g |%6g|\n", 12.3, 34.5) // |  12.3 |  34.5|
	fmt.Printf("|%6f |%6f|\n", 12.3, 34.5) // |12.300000 |34.500000|

	// floating with post deciaml 2 digits only
	fmt.Printf("|%6.2f |%6.2f|\n", 12.3, 34.5)   // | 12.30 | 34.50|
	fmt.Printf("|%-6.2f |%-6.2f|\n", 12.3, 34.5) // |12.30  |34.50 |

	fmt.Printf("|%-06.2f |%-06.2f|\n", 12.3, 34.5) // |12.30  |34.50 |
	fmt.Printf("|%06.2f |%06.2f|\n", 12.3, 34.5)   // |012.30 |034.50 |
	fmt.Println()

	fmt.Printf("|%6s|%6s|\n", "foo", "b")   // |   foo|     b|
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // |foo   |b     |
	fmt.Println()

	// creating new string from formatting
	myStr := fmt.Sprintf("A %s baby was born %s\n", "beutiful", "today")
	fmt.Println(myStr) // A beutiful baby was born today

	fmt.Fprintf(os.Stdout, "A %s baby was born %s\n", "beutiful", "today")
	fmt.Fprintf(os.Stderr, "A %s baby was born %s\n", "beutiful", "today")
	fmt.Println()

	// Struct formatting
	type point struct {
		x, y int
	}
	p1 := point{1, 2}
	fmt.Println("p1 = ", p1) // p1 =  {1 2}
	fmt.Printf("%v\n", p1)   // {1 2}
	fmt.Printf("%+v\n", p1)  // {x:1 y:2}
	fmt.Printf("%-v\n", p1)  // {1 2}
	fmt.Printf("%#v\n", p1)  // main.point{x:1, y:2}
	fmt.Println()

	fmt.Println("reflect.TypeOf(p):", reflect.TypeOf(p1)) //  main.point
	fmt.Printf("%T\n", p1)                                // main.point
	fmt.Printf("%t\n", true)
}
