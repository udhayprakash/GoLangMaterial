package main

/*
Purpose: MX Records identify the servers that can exchange emails.

*/
import (
	"fmt"
	"net"
)

func main() {
	mxRecords, _ := net.LookupMX("google.com")
	for _, mx := range mxRecords {
		fmt.Println(mx.Host, mx.Pref)
	}
}
