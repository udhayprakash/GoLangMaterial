package main

import (
	"fmt"
	"net"
)

func main() {
	for _, service := range []string{
		"http",
		"https",
		"smtp",
		"tls",
		"ssl",
		"invalidService",
	} {
		port, err := net.LookupPort("tcp", service)
		if err == nil {
			fmt.Printf("port for %s is %v\n", service, port)
		} else {
			fmt.Println(err)
		}
	}

}
