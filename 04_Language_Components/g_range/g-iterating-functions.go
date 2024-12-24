package main

import "fmt"

// Iterator design pattern

// Custom iterator function for a slice
func sliceIterator(slice []int) func() (int, bool) {
    i := 0
    return func() (int, bool) {
        if i >= len(slice) {
            return 0, false // No more elements
        }
        val := slice[i]
        i++
        return val, true // Return current value and true if more elements exist
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}

    // Initialize the iterator
    iter := sliceIterator(numbers)

    // Manually iterate using the iterator function
    for {
        num, ok := iter()
        if !ok {
            break // Exit the loop when there are no more elements
        }
        fmt.Println(num)
    }
}
