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
	const (
		a1 = 10
		a2 = 11
		a3 = 12
		a4 = 13
	)
	fmt.Println(a1, a2, a3, a4) // 10 11 12 13

	/*
		Go does not have enumerate types

		Instead, you can use the special name iota in a single
		const declaration to get a series of increasing values.
	*/
	const (
		b1 = iota
		b2 = iota
		b3 = iota
	)
	fmt.Println(b1, b2, b3) // 0 1 2

	const (
		c1 = 10 + iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3) // 10 1 2

	const (
		d1 = 10 + iota
		d2
		d3
	)
	fmt.Println(d1, d2, d3) // 10 11 12

	const (
		e1 = 10
		e2
		e3
	)
	fmt.Println(e1, e2, e3) // 10 10 10

	// ---------COUNTING BACKWARD----------------
	const (
		f1 = (10 - iota) // 10 - 0 = 10
		f2               // 10 -1  = 9
		f3               // 10 -2 = 8
	)
	fmt.Println(f1, f2, f3) // 10 9 8

	const (
		g1 = (0 - iota) // 0 - 0 = 0
		g2
		g3
	)
	fmt.Println(g1, g2, g3) // 0 -1 -2

	const (
		h1 = iota + 3
		_  // blank identifier to skip a value in a list of constants.
		h2
		h3
	)
	fmt.Println(h1, h2, h3) // 3 5 6

	// ---------------------------------------
	const (
		January int = 1 + iota
		February
		March
		April int = 10 //+ iota  // iota index value will reset
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
		q1, q2, q3 = iota, iota, iota
		r1, r2, r3
		s1, s2, s3
	)

	fmt.Printf("q1: %d, q: %d, q3: %d r1: %d, r2: %d, r3: %d s1: %d, s2: %d, s3: %d", q1, q2, q3, r1, r2, r3, s1, s2, s3)

	// multi-line strings - `` quotes
	fmt.Printf(`
	q1: %d, q: %d, q3: %d 
	r1: %d, r2: %d, r3: %d 
	s1: %d, s2: %d, s3: %d`, q1, q2, q3, r1, r2, r3, s1, s2, s3)

}
