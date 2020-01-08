package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)

	log.Output(1,"One")
	log.Output(2,"Two")
	log.Output(3,"Three")
	log.Output(4,"Four")
}

// Output writes the output for a logging event.
// The string s contains the text to print after the prefix
// specified by the flags of the Logger. A newline is
// appended if the last character of s is not already a newline.

// Calldepth is the count of the number of frames to skip when
// computing the file name and line number if Llongfile or
// Lshortfile is set; a value of 1 will print the details for
// the caller of Output.
