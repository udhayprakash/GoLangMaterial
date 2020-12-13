package main

import "fmt"

/*
KEYWORDS in Go
----------------
break 		default 	func 	interface 	select
case 		defer 		go 		map 		struct
chan 		else 		goto 	package 	switch
const 		fallthrough if 		range 		type
continue 	for 		import 	return 		var

	In addition, three dozen predeclared names.
Constants : true 	false 	iota 	nil
Types     : int 	int8 	int32 	int64

unit 	unit8 	unit16 		uint32 		uint64 	uintptr
float32 float64 complex128 	complex64
bool 	byte 	rune 		string 		error
*/

func main() {
	// NOTE 1: keywords should not be used as identifiers
	// break := "one" //syntax error: unexpected := at end of statement
	// true := "one" //: true declared but not used
	
	break_one := "one"
	fmt.Println("break_one =", break_one)

	// NOTE 2: CamelCase is recommended for identifiers

	breakOne := "one"
	fmt.Println("breakOne=", breakOne)

	costOfMangos := 34
	fmt.Println("costOfMangos=", costOfMangos)

	NoOfProcessesRunning := 3
	fmt.Println("NoOfProcessesRunning=", NoOfProcessesRunning)
}
