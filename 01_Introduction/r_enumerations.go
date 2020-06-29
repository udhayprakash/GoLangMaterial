/*
Purpose: Enum
	- Grouping and expecting only some related values
	- Sharing common behavior
	- Avoids using invalid values
	- To increase the code readability and the maintainability

*/
package main

import "fmt"

func main() {
	// Go does not have enumerate types
	// Instead, you can use the special name iota in a single const declaration to get a series of increasing values.
	// When an initialization expression is omitted for a const, it reuses the preceding expression.
	const (
		C0 = iota
		C1 = iota
		C2 = iota
	)
	fmt.Println(C0, C1, C2) // "0 1 2"
	fmt.Println()

	// ---------------------------------------
	const (
		red   = iota // red == 0
		blue         // blue == 1
		green        // green == 2
	)
	fmt.Println("red   =", red)
	fmt.Println("blue  =", blue)
	fmt.Println("green =", green)
	fmt.Println()

	// ---------COUNTING BACKWARD----------------
	const (
		max = 10
	)
	const (
		a = (max - iota) // 10
		b                // 9
		c                // 8
	)
	fmt.Println(a, b, c) // "1 2 4"

	// ---------------------------------------
	const (
		p1 = iota + 1
		p3
		_ // blank identifier to skip a value in a list of constants.
		p4
	)
	fmt.Println(p1, p3, p4) // "1 2 4"
	fmt.Println()

	// ---------------------------------------
	const (
		January int = 1 + iota
		February
		March
		April int = 10 //+ iota
		May
	)
	fmt.Println("January  :", January)
	fmt.Println("February :", February)
	fmt.Println("March    :", March)
	fmt.Println("April    :", April)
	fmt.Println("May      :", May)

	// ---------- Multiple iotas in a single line
	const (
		// Active = 0, Moving = 0, Running = 0
		Active, Moving, Running = iota, iota, iota

		// Passive = 1, Stopped = 1, Stale = 1
		Passive, Stopped, Stale
	)
	fmt.Print("Active, Moving, Running :")
	fmt.Println(Active, Moving, Running)
	fmt.Print("Passive, Stopped, Stale :")
	fmt.Println(Passive, Stopped, Stale)

}
