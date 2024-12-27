package main

import "fmt"

/*
When array elements of an array are arrays,
then it’s called a multi-dimensional array.
*/
func main() {
	oneDArray := [3]int{11, 22, 33}
	twoDArray := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{4, 5},
	}
	twoDArray2 := [3][3]string{
		{"Hyderbad", "Amaravathi", "Chennai"},
		{"Kochi", "Banglore", "Panagi"},
		{"Mumbai", "Ahmedabad", "Jaipur"},
	}

	fmt.Printf(`
		oneDArray = %#v
		twoDArray = %#v
		twoDArray2 = %#v
	`, oneDArray, twoDArray, twoDArray2)

	a2 := [...][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		{5, 6},
		{7, 8},
	}
	fmt.Println("a2:", a2)
	fmt.Println()

	// Iterating multi-dimensional arrays
	for index1, childArray := range a2 {
		fmt.Println("At index ", index1, " we have", childArray)
		for _, element := range childArray {
			fmt.Printf("\telement: %d\n", element)
		}
	}

	// ===================
	anArray := [4]int{1, 2, 4, -4}
	twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}

	fmt.Println("The length of", anArray, "is", len(anArray))
	fmt.Println("The first element of", twoD, "is", twoD[0][0])
	fmt.Println("The length of", threeD, "is", len(threeD))

	for i := 0; i < len(threeD); i++ {
		v := threeD[i]
		for j := 0; j < len(v); j++ {
			m := v[j]
			for k := 0; k < len(m); k++ {
				fmt.Print(m[k], " ")
			}
		}
		fmt.Println()
	}

	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
		fmt.Println()
	}
}
