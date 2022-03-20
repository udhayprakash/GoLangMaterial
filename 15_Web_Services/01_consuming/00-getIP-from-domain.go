package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	for _, website := range []string{
		"google.com",
		"aha.com",
		"netflix.com",
		"facebook.com",
		"instagram.com",
		"amazon.com",
		"zee5.com",
		"whatsapp.com",
		"twitter.com",
	} {
		ip_list, err := net.LookupIP(website)
		if err == nil {
			fmt.Println(website)
			for _, ip := range ip_list {
				fmt.Println("\t", ip)
			}
			cName, err := net.LookupCNAME(website)
			if err == nil {
				fmt.Println("\t", cName)
			} else {
				fmt.Print("Lookup CNAME failed with err: ", err)
			}
			DNSTextRecords, err := net.LookupTXT(website)
			if err == nil {
				fmt.Println("DNSTextRecords:", DNSTextRecords)
			}

		} else {
			log.Fatal("IP lookup failed. Error is: ", err)
		}
	}
}
