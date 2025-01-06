package main

import (
	"fmt"
	"log"
)

/*
Ldate         = 1 << iota // 1: the date in the local time zone (e.g., 2009/01/23)
Ltime                     // 2: the time in the local time zone (e.g., 01:23:23)
Lmicroseconds             // 4: microsecond resolution (e.g., 01:23:23.123123)
Llongfile                 // 8: full file name and line number (e.g., /a/b/c/d.go:23)
Lshortfile                // 16: final file name element and line number (e.g., d.go:23)
LUTC                      // 32: use UTC instead of local time zone
Lmsgprefix                // 64: move the prefix to before the message
LstdFlags = Ldate | Ltime // 3: initial values for the standard logger
*/
func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}
func main() {
	flags := log.Flags()
	log.Println("Flags:", flags)
	// returns an integer representing the current logging flags set for the standard logger
	// 1 (log.Ldate) + 4 (log.Lmicroseconds) + 16 (log.Lshortfile) = 21

	// Decode the flags
	if flags&log.Ldate != 0 {
		fmt.Println("Ldate is set")
	}
	if flags&log.Ltime != 0 {
		fmt.Println("Ltime is set")
	}
	if flags&log.Lmicroseconds != 0 {
		fmt.Println("Lmicroseconds is set")
	}
	if flags&log.Llongfile != 0 {
		fmt.Println("Llongfile is set")
	}
	if flags&log.Lshortfile != 0 {
		fmt.Println("Lshortfile is set")
	}
	if flags&log.LUTC != 0 {
		fmt.Println("LUTC is set")
	}
	if flags&log.Lmsgprefix != 0 {
		fmt.Println("Lmsgprefix is set")
	}
}
