package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
Problem :
You have a function that will load a resource such as reading configuration data from an external file as part of the initialization process. However, each time when the main program calls the function, it will read the configuration data again and again.
How to ensure that certain part of the function only execute once only? Load once instead of reading the configuration data again and again?

Solution :
Use the sync.Once object. Below is an example that demonstrates how sync.Once object can be used to wrap some part of the code inside a function to only execute once.
*/

const configurationFile = "servername = 123.google.com"

var onlyOnce sync.Once
var config strings.Reader

func loadConfigurations() {

	// we only want to load the configuration once
	onlyOnce.Do(func() {
		fmt.Println("Load run-time configuration first and the only time. ")
		config = *strings.NewReader(configurationFile)
	})

	fmt.Println(config)
}

func main() {
	loadConfigurations()
	go loadConfigurations()
	go loadConfigurations()
	go loadConfigurations()
	loadConfigurations()

	// keep running until user press Control-C
	for {
	}
}
