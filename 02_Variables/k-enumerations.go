package main

import "fmt"

/*
Purpose: Enum
	- Grouping and expecting only some related values
	- Sharing common behavior
	- Avoids using invalid values
	- To increase the code readability and the maintainability

*/
func main() {
	// const (
	// 	a1 = 10
	// 	a2 = 11
	// 	a3 = 12
	// 	a4 = 13
	// )

	/*
		Go does not have enumerate types

		Instead, you can use the special name iota in a single
		const declaration to get a series of increasing values.
	*/
	const (
		C0 = iota
		C1 = iota
		C2 = iota
	)
	fmt.Println(C0, C1, C2) // "0 1 2"
	fmt.Println()

	// ---------------------------------------
	// When an initialization expression is omitted for a const,
	// 	it reuses the preceding expression.
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
	fmt.Println(a, b, c) // 10 9 8

	// ---------------------------------------
	const (
		p1 = iota + 3
		_  // blank identifier to skip a value in a list of constants.
		p3
		p4
	)
	fmt.Println(p1, p3, p4) // 3 5 6
	fmt.Println()

	// ---------------------------------------
	const (
		January int = 1 + iota
		February
		March
		April int = 10 //+ iota
		May
		June
	)
	fmt.Println("January  :", January)  // 1
	fmt.Println("February :", February) // 2
	fmt.Println("March    :", March)    // 3
	fmt.Println("April    :", April)    // 10
	fmt.Println("May      :", May)      // 10
	fmt.Println("June     :", June)     // 10

	// ---------- Multiple iotas in a single line
	const (
		// Active = 0, Moving = 0, Running = 0
		Active, Moving, Running = iota, iota, iota

		// Passive = 1, Stopped = 1, Stale = 1
		Passive, Stopped, Stale
		something1, something2, something3
	)
	fmt.Print("Active, Moving, Running :")
	fmt.Println(Active, Moving, Running)
	fmt.Print("Passive, Stopped, Stale :")
	fmt.Println(Passive, Stopped, Stale)
	fmt.Println("something1, something2, something3:", something1, something2, something3)
}
