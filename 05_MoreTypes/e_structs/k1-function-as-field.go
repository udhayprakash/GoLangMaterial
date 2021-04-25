// Go program to illustrate the function
// as a field in Go structure
// Using anonymous function
package main

import "fmt"

// Creating structure
type Author struct {
	name      string
	language  string
	Tarticles int
	Particles int
	Pending   func(int, int) int
}

// Main method
func main() {

	// Initializing the fields
	// of the structure
	result := Author{
		name:      "Sonia",
		language:  "Java",
		Tarticles: 340,
		Particles: 259,
		Pending: func(Ta int, Pa int) int {
			return Ta - Pa
		},
	}

	// Display values
	fmt.Println("Author's Name: ", result.name)
	fmt.Println("Language: ", result.language)
	fmt.Println("Total number of articles: ", result.Tarticles)

	fmt.Println("Total number of published articles: ",
		result.Particles)

	fmt.Println("Pending articles: ", result.Pending(result.Tarticles,
		result.Particles))
}
