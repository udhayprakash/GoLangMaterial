package main

import (
	"fmt"
	"strings"
)

func main() {

	hostings := map[int]string{0: "digitalocean.com",
		1: "linode.com",
		2: "azure.com",
		3: "aws.amazon.com",
		4: "heroku.com"}

	fmt.Println("Before filtering :")
	fmt.Println("-------------------------")
	for number, company := range hostings {
		fmt.Println("Number : ", number, "Company : ", company)
	}

	fmt.Println("\n\nFilter by index number. Need to know index number in advance :")
	fmt.Println("-------------------------")
	companyName, ok := hostings[3] // select index number 3

	if ok {
		fmt.Println("Hosting company name : ", companyName)
	} else {
		fmt.Println(companyName + " not found")
	}

	// filter aws.amazon.com from the hostings map

	fmt.Println("\n\nFilter map elements by name :")
	fmt.Println("-------------------------")
	for number, company := range hostings {
		if strings.EqualFold(strings.ToLower(company), "aws.amazon.com") {
			fmt.Printf("[%d] : %s\n", number, company)
		}
	}

	// filter company names that contain the character "u" from the hostings map
	fmt.Println("\n\nFilter map elements by a single character :")
	fmt.Println("-------------------------")
	for number, company := range hostings {
		if strings.Contains(strings.ToLower(company), "u") {
			fmt.Printf("[%d] : %s\n", number, company)
		}
	}

}
