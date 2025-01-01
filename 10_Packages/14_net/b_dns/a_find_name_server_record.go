package main

import (
	"fmt"
	"net"
)

/*
Purpose: Name Server(NS) records describe the authorized name servers for the zone.
*/
func main() {
	for _, eachWebSite := range []string{"google.com", "linkedin.com", "twitter.com"} {
		nameServer, _ := net.LookupNS(eachWebSite)
		for _, ns := range nameServer {
			fmt.Println("eachWebSite:", eachWebSite, "  ns:", ns)
		}
		fmt.Println()
	}
}
