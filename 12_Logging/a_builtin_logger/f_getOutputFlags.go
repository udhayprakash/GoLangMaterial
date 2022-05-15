package main

import "log"


func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}
func main() {
	flags := log.Flags()
	log.Println("Flags:", flags)
}
