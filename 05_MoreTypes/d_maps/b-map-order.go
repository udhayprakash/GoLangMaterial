package main

import "fmt"

func main() {
	tagsViews := map[string]int{
		"unix":   0,
		"python": 1,
		"go":     2,
		"golang": 3,
		"devops": 4,
		"gc":     5,
	}
	fmt.Println("tagsViews = ", tagsViews) // displays in key based ascending order
	//  map[devops:4 gc:5 go:2 golang:3 python:1 unix:0]

	for range tagsViews {
		fmt.Println("simple loop")
	}

	for key := range tagsViews {
		fmt.Println("key=", key)
	}

	for key, value := range tagsViews {
		fmt.Println("There are", value, "views for", key)
	}
}

// NOTE: Map order is undefined
