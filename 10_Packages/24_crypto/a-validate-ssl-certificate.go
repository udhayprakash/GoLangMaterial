package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	sites := [...]string{
		"facebook.com:80",
		"facebook.com:443",
		"google.com:80",
		"google.com:443",
		"freecodecamp.org:80",
		"freecodecamp.org:443",
	}
	for _, site := range sites {
		conn, err := tls.Dial("tcp", site, nil)
		if err != nil {
			fmt.Println("Server doesn't support SSL certificate err: " + err.Error())
			continue
		}
		fmt.Println(site, "==>", conn.LocalAddr())

		// Check whether SSL certificatre and website hostname match
		err = conn.VerifyHostname(site)
		if err != nil {
			fmt.Println("Hostname doesn't match with certificate: " + err.Error())
		}

	}
}