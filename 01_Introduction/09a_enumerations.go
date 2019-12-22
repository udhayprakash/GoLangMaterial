package main

import "fmt"

func main() {
	// Go does not hve enumerate types
	//Instead, you can use the special name iota in a single const declaration to get a series of increasing values.
	// When an initialization expression is omitted for a const, it reuses the preceding expression.

	const (
		red   = iota // red == 0
		blue         // blue == 1
		green        // green == 2
	)

	fmt.Println("red   =", red)
	fmt.Println("blue  =", blue)
	fmt.Println("green =", green)

	const (
		January int = 1 + iota
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)
	fmt.Println("January  :", January)
	fmt.Println("February :", February)
	fmt.Println("March    :", March)
	fmt.Println("April    :", April)
	fmt.Println("May      :", May)
	fmt.Println("June     :", June)
	fmt.Println("July     :", July)
	fmt.Println("August   :", August)
	fmt.Println("September:", September)
	fmt.Println("October  :", October)
	fmt.Println("November :", November)
	fmt.Println("December :", December)
}
