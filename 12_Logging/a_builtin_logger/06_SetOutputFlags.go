package main

import (
	"log"
)

func main() {
	log.Println("Default")

	// the date in the local time zone: 2009/01/23
	log.SetFlags(log.Ldate)
	log.Println("Using log.Ldate")

	// the time in the local time zone: 01:23:23
	log.SetFlags(log.Ltime)
	log.Println("Using log.Ltime")

	// microsecond resolution: 01:23:23.123123
	log.SetFlags(log.Lmicroseconds)
	log.Println("Using log.Lmicroseconds")

	// full file name and line number: /a/b/c/d.go:23
	log.SetFlags(log.Llongfile)
	log.Println("Using log.Llongfile")

	// final file name element and line number: d.go:23
	log.SetFlags(log.Lshortfile)
	log.Println("Using log.Lshortfile")

	// if Ldate or Ltime is set, use UTC rather than the local time zone
	log.SetFlags(log.Ltime | log.LUTC)
	log.Println("Using log.LUTC")

	// initial values for the standard logger
	log.SetFlags(log.LstdFlags)
	log.Println("Using log.LstdFlags")
}