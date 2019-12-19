package main

/*

Maps are Go’s built-in associative data type
(sometimes called hashes or dicts in other languages).

*/
import "fmt"
func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map         :", m)	// map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println("v1          :", v1)	// 7

	fmt.Println("len         :", len(m))	// 2

	delete(m, "k2")
	fmt.Println("map         :", m) // map[k1:7]

	val, isKeyPresent := m["k1"]
	fmt.Println("\nval         :", val)
	fmt.Println("isKeyPresent:", isKeyPresent)

	val, isKeyPresent = m["k6"]
	fmt.Println("\nval         :", val)
	fmt.Println("isKeyPresent:", isKeyPresent)

	//blank identifier _ -used when we dont need a value
	_, isKeyPresent = m["k8"]

	// Declaring and initializing new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map         :", n)
}
