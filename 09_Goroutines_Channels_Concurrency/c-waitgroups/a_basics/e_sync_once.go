package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
Sync.One  will run the registered inline function only one, for entire program execution

Mostly used in one time execution scenarios like loading config data, etc
*/

const configurationFile = `
	servername=https://photos.google.com
	verify=True
	LanguageDefault:Golang
`

var onlyOnce sync.Once
var config strings.Reader

func loadConfigurations() {

	// we only want to load the configuration once
	onlyOnce.Do(func() {
		fmt.Println("Load run-time configuration first and the only time. ")
		config = *strings.NewReader(configurationFile)
	})

	fmt.Println("\nDisplaying Loaded config data:", config)
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
