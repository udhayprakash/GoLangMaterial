// LZW Data Compression in Golang
package main

import "fmt"

// LZW is a data compression method that takes advantage of this repetition. This compression method, (1) start with an initial model, (2) read data piece by piece, (3) and update the model and encode the data as you go along. LZW is a "dictionary"-based compression algorithm. LZW encodes data by referencing a dictionary. To encode a sub-string, only a single code number, corresponding to that sub-string's index in the dictionary, needs to be written to the output file.
func compressLZW(testStr string) []int {
	code := 256
	dictionary := make(map[string]int)
	for i := 0; i < 256; i++ {
		dictionary[string(i)] = i
	}

	currChar := ""
	result := make([]int, 0)
	for _, c := range []byte(testStr) {
		phrase := currChar + string(c)
		if _, isTrue := dictionary[phrase]; isTrue {
			currChar = phrase
		} else {
			result = append(result, dictionary[currChar])
			dictionary[phrase] = code
			code++
			currChar = string(c)
		}
	}
	if currChar != "" {
		result = append(result, dictionary[currChar])
	}
	return result
}

func decompressLZW(compressed []int) string {
	code := 256
	dictionary := make(map[int]string)
	for i := 0; i < 256; i++ {
		dictionary[i] = string(i)
	}

	currChar := string(compressed[0])
	result := currChar
	for _, element := range compressed[1:] {
		var word string
		if x, ok := dictionary[element]; ok {
			word = x
		} else if element == code {
			word = currChar + currChar[:1]
		} else {
			panic(fmt.Sprintf("Bad compressed element: %d", element))
		}

		result += word

		dictionary[code] = currChar + word[:1]
		code++

		currChar = word
	}
	return result
}

func main() {
	fmt.Print("Enter any string :")
	var testStr string
	fmt.Scanln(&testStr)

	compressed := compressLZW(testStr)
	fmt.Println("\nAfter Compression :", compressed)

	uncompression := decompressLZW(compressed)
	fmt.Println("\nAfter Uncompression :", uncompression)
}
