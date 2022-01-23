package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Please input at least 4 integers, separated by spaces: ")
	if scanner.Scan() {
		initialArray := strings.Fields(scanner.Text())
		// convert array of string to array of ints
		arrayOfInts := ConvertStringsToInts(initialArray)
		// divide array into four slices of equal length
		splitArray := SplitInChunks(arrayOfInts, 4)
		channel := make(chan []int)

		var result []int

		for i := 0; i < 4; i++ {
			go SortChunk(splitArray[i], channel)
			result = append(result, <-channel...)
		}

		sort.Ints(result)
		fmt.Printf("Resulting array: %v", result)

	} else {
		fmt.Printf("Error occurred scanning user input: %s", scanner.Err())
	}
}

func SplitInChunks(original []int, numberOfChunks int) [][]int {
	var split [][]int

	chunkSize := (len(original) + numberOfChunks - 1) / numberOfChunks

	for i := 0; i < len(original); i += chunkSize {
		end := i + chunkSize

		if end > len(original) {
			end = len(original)
		}

		split = append(split, original[i:end])
	}

	return split
}

func SortChunk(chunk []int, channel chan []int) {

	fmt.Printf("Sorting subarray %v...\n", chunk)
	sort.Ints(chunk)
	channel <- chunk
}

func ConvertStringsToInts(input []string) []int {

	var outs []int

	for _, i := range input {
		intFromString, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		outs = append(outs, intFromString)
	}

	return outs
}
