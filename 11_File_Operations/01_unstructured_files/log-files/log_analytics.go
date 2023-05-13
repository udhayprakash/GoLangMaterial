package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func prettyPrintMap(m map[string]int) {
	fmt.Println("--------")
	fmt.Println("KEY\tVALUE")
	fmt.Println("--------")
	for k, v := range m {
		fmt.Printf("%-8s %d\n", k+":", v)
	}
	fmt.Println("--------")
}

func main() {
	// Open the access.log file for reading
	file, err := os.Open("access.log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a map to store the count of each URL
	urlCount := make(map[string]int)

	// Read the file line by line and count the URLs
	scanner := bufio.NewScanner(file)
	mostAccessedCount := 0
	mostAccessedURL := ""
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) <= 5 {
			continue
		}
		url := strings.Split(fields[6], "?")[0]

		urlCount[url]++
		if urlCount[url] > mostAccessedCount {
			mostAccessedURL = url
			mostAccessedCount = urlCount[url]
		}
	}

	prettyPrintMap(urlCount)

	// Print the result
	fmt.Printf("The most accessed URL is %s with %d hits.\n",
		mostAccessedURL, mostAccessedCount)
}
