package main
/*
Variables declared without an explicit initial value
are given their zero value.

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.

*/

import "fmt"

func main() {
	var i int        // 0
	var f float64    // 0
	var b bool       // false
	var s string     // ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

/*
ZERO VALUES:
~~~~~~~~~~~
	var s string // defaults to ""
	var r rune   // defaults to 0
	var by byte  // defaults to 0
	var in int   // defaults to 0
	var ui uint  // defaults to 0
	var f float32 // defaults to 0
	var c complex64 // defaults to 0
	var b bool // defaults to false

	var arr [2]int // defaults to [0 0]
	var obj struct {
		a bool
		b [2]int
	}{}            // defaults to {false [0 0]}
	
	var si []int // defaults to nil | []
	var ch chan string // defaults to nil
	var m map[string]int  // defaults to nil | map[]
	var fn func() // defaults to nil
	var itf chan interface{} // defaults to nil
	var ptr *string // defaults to nil

*/
