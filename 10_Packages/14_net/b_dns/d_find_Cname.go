package main

/*
Purpose: CNAME
		- Canonical Name
		- It is domain and sub-domain text aliases to bind traffic.

*/

import (
	"fmt"
	"net"
)

func main() {
	cName, _ := net.LookupCNAME("m.facebook.com")
	fmt.Println("cName:", cName)
}
