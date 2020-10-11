package main

import "fmt"

func main(){
	// Declaring and initializing new map in the same line
	n := map[string]string{"a": "apple", "b": "ball", "c": "cat"}
	fmt.Println("map         :", n)

	fmt.Println(n["a"]) // apple
	fmt.Println(n["a1"]) // gives nothing when the key is not defined in that map

	if x, ok := n["a"]; ok {
		fmt.Println("m['a'] = ", x) // "apple"
	}

	if x, ok := n["a1"]; ok {
		fmt.Println("m['a1'] = ", x)
	}

	// To initialize a map with some data, use a map literal:
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	fmt.Println(commits)

}

/*
Assignment:
	Using random module, you can toss a coin to give either head or tail.
	After 100 trails, give how many times, head is greater/lower than the tails.

*/