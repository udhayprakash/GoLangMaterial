package main

import "log"

func main() {

	// Enable line numbers in logging. By default, it is disabled
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Will print: "[date] [time] [filename]:[line]: [text]"
	log.Println("Logging w/ line numbers on golangcode.com")
}