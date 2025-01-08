package main

import (
	"bufio"
	"fmt"
	"strings"
)

const singleLine string = "I'd love to have some coffee right about now"
const multiLine string = "Reading is my...\r\n favourite"

func main() {
	// fmt.Println("Lenght of singleLine input is ", len(singleLine))
	// fmt.Println("Lenght of singleLine input is " + strconv.Itoa(len(singleLine)))
	fmt.Printf("Lenght of singleLine input is %d\n", len(singleLine))

	str := strings.NewReader(singleLine)

	br := bufio.NewReaderSize(str, 25)

	fmt.Println("\n---Peek---")
	// Peek - Case 1: Simple peek implementation
	b, err := br.Peek(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", b) // output: "I'd"

	// Peek - Case 2: Peek larger than buffer size
	b, err = br.Peek(30) // 3 to 33
	if err != nil {
		fmt.Println(err) // output: "bufio: buffer full"
	}

	// Peek - Case 3: Buffer size larger than string
	br_large := bufio.NewReaderSize(str, 50)
	b, err = br_large.Peek(50)
	if err != nil {
		fmt.Println(err) // output: EOF
	}

	// Peek - Case 4: Buffer size exact as  string
	br_exact := bufio.NewReaderSize(str, len(singleLine))
	b, err = br_exact.Peek(len(singleLine))
	if err != nil {
		fmt.Println(err) // output: EOF
	}

	// ReadSlice
	fmt.Println("\n---ReadSlice---")
	str = strings.NewReader(multiLine)
	r := bufio.NewReader(str)
	for {
		token, err := r.ReadSlice('.')
		if len(token) > 0 {
			fmt.Printf("Token (ReadSlice): %q\n", token)
		}
		if err != nil {
			break
		}
	}

	// ReadLine
	fmt.Println("\n---ReadLine---")
	str = strings.NewReader(multiLine)
	r = bufio.NewReader(str)
	for {
		token, _, err := r.ReadLine()
		if len(token) > 0 {
			fmt.Printf("Token (ReadLine): %q\n", token)
		}
		if err != nil {
			break
		}
	}

	// ReadBytes
	fmt.Println("\n---ReadBytes---")
	str = strings.NewReader(multiLine)
	r.Reset(str)
	for {
		token, err := r.ReadBytes('\n')
		fmt.Printf("Token (ReadBytes): %q\n", token)
		if err != nil {
			break
		}
	}

	// Scanner
	fmt.Println("\n---Scanner---")
	str = strings.NewReader(multiLine)
	scanner := bufio.NewScanner(str)
	for scanner.Scan() {
		fmt.Printf("Token (Scanner): %q\n", scanner.Text())
	}

}
