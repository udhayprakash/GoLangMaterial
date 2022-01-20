package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {

	// Explicit type conversion.
	fmt.Println(int64(d), reflect.TypeOf(int64(d)))
	fmt.Println()
	
	f := 5.1
	i := int(f) // convert float to int

	fmt.Println(f, reflect.TypeOf(f))
	fmt.Println(i, reflect.TypeOf(i))

	// Integers ===============================================================
	var a uint16 = 0x10fe //  bit pattern: 0001 0000 1111 1110

	// Truncation
	b := int8(a) // -2        bit pattern:           1111 1110

	// Sign extension
	c := uint16(b) // 0xfffe  bit pattern: 1111 1111 1111 1110
	fmt.Println(c)

	// Floats ================================================================
	var f1 float64 = 1.9
	n := int64(f1) // 1
	n = int64(-f1) // -1

	n = 1234567890
	g := float32(n) // 1.234568e+09
	fmt.Println(g)

	// Integer to String ====================================================
	fmt.Println("string(97) =", string(97)) // "a"
	fmt.Println("string(-1) =", string(-1)) // "\ufffd" == "\xef\xbf\xbd"
	fmt.Println()

	// To get the decimal string representation of an integer:
	fmt.Println("strconv.Itoa(97)=", strconv.Itoa(97)) // "97"

	//fmt.Println(int("55"))
	fmt.Println(int('5'))

	// Strings and byte slices ===============================================
	// Converting a slice of bytes to a string type yields a string whose successive bytes are the elements of the slice.
	fmt.Println("string([]byte{97, 230, 151, 165})", string([]byte{97, 230, 151, 165})) // "a日"
	// Converting a value of a string type to a slice of bytes type yields a slice whose successive elements are the bytes of the string.
	fmt.Println(`[]byte("a日")`, []byte("a日")) // []byte{97, 230, 151, 165}

	// Strings and rune slices ===============================================
	// Converting a slice of runes to a string type yields a string that is the concatenation of the individual rune values converted to strings.
	fmt.Println("string([]rune{97, 26085}):", string([]rune{97, 26085})) // "a日"
	// Converting a value of a string type to a slice of runes type yields a slice containing the individual Unicode code points of the string.
	fmt.Println(`[]rune("a日")            :`, []rune("a日")) // []rune{97, 26085}

	// Underlying type =======================================================
	// A non-constant value can be converted to type T if it has the same underlying type as T.
	type (
		T1 int64
		T2 T1
	) // the underlying type of int64, T1, and T2 is int64.

	var n1 int64 = 12345
	fmt.Println(n1)                // 12345
	fmt.Println(time.Duration(n1)) // 12.345µs

	// Implicit conversions ==================================================
	// The only implicit conversion in Go is when an untyped constant is used in a situation where a type is required.
	var f2 float64
	f2 = 1                // Same as: f2 = float64(1)
	t2 := 2 * time.Second // Same as: t2 := time.Duration(2) * time.Second

	fmt.Println(f2, t2)

	n3 := 1   // Same as: n3 := int(1)
	f3 := 1.0 // Same as: f3 := float64(1.0)
	s3 := "A" // Same as: s3 := string("A")
	c3 := 'A' // Same as: c3 := rune('A')

	// Illegal implicit conversions are caught by the compiler:
	//var b3 byte = 256 // Same as: var b3 byte = byte(256) // constant 256 overflows byte

	fmt.Println(n3, f3, s3, c3)

	// Pointers ==============================================================
	// The Go compiler does not allow conversions between pointers and integers.
	// The package unsafe implements this functionality under restricted circumstances.
	// But, used only in LOW-LEVEL programming
}
