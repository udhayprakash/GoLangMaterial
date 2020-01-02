// [_Command-line flags_](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// are a common way to specify options for command-line
// programs. For example, in `wc -l` the `-l` is a
// command-line flag.
package main

import (
	"flag"
	"fmt"
)

// Param 1: Argument name, called like ./args --taska
// Param 2: Sets the default state, false would mean task B runs by default
// Param 3: Is the description of that command, shown if you call ./args --help
var toRunTaskA = flag.Bool("taska", false, "Whether to run task A or taskB")

func main() {

	flag.Parse()

	if *toRunTaskA {
		fmt.Println("running Task A ...")
	} else {
		fmt.Println("running Task B ...")
	}
}
