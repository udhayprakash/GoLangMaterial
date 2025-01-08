package main

import (
	"fmt"
	"os"
	"testing"
)

// set function go initialize resources
func setup() {
	fmt.Println("Creating temporary file...")
	file, err := os.CreateTemp("", "example.txt")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer file.Close()
	fmt.Println("Temporary file created:", file.Name())
}

// teardown function to clean up resources
func teardown() {
	fmt.Println("Removing temporary file...")
	// Assuming the temporary file name is known or stored somewhere
	err := os.Remove("example.txt")
	if err != nil {
		fmt.Println("Error removing temporary file:", err)
	} else {
		fmt.Println("Temporary file removed")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Setting up resources...")
	// Setup code here
	setup()

	// Run tests
	code := m.Run()

	fmt.Println("Tearing down resources...")
	// Teardown code here
	teardown()
	os.Exit(code) // 0 for success, non-zero for failure
}

func TestSomething(t *testing.T) {
	t.Log("Testing Something Case")

}
