package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

// Reading twice from same io.Reader
func main() {
	originalReader := strings.NewReader("test string in test reader")

	b, err := ioutil.ReadAll(originalReader)
	if err != nil {
		panic(err)
	}

	reader1 := bytes.NewReader(b)
	b1, err := ioutil.ReadAll(reader1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b1))

	reader2 := bytes.NewReader(b)
	b2, err := ioutil.ReadAll(reader2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b2))
}
