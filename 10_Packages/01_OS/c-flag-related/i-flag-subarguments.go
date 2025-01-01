package main

import (
	"flag"
	"fmt"
	"os"
)

/*
Purpose: Some command-line tools, like the go tool or git
have many subcommands, each with its own set of flags.
For example, go build and go get are two different
subcommands of the go tool. The flag package lets us
easily define simple subcommands that have their own flags.
*/

func main() {

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

/*
$ go run i-flag-subarguments.go 
expected 'foo' or 'bar' subcommands
exit status 1
$ go run i-flag-subarguments.go -h
expected 'foo' or 'bar' subcommands
exit status 1
$ go run i-flag-subarguments.go foo
subcommand 'foo'
  enable: false
  name: 
  tail: []
$ go run i-flag-subarguments.go -foo
expected 'foo' or 'bar' subcommands
exit status 1
$ go run i-flag-subarguments.go foo -enable 
subcommand 'foo'
  enable: true
  name: 
  tail: []
$ go run i-flag-subarguments.go foo -enable true
subcommand 'foo'
  enable: true
  name: 
  tail: [true]
$ go run i-flag-subarguments.go foo -enable false
subcommand 'foo'
  enable: true
  name: 
  tail: [false]
$ go run i-flag-subarguments.go foo -enable false name=ravi
subcommand 'foo'
  enable: true
  name: 
  tail: [false name=ravi]
$ go run i-flag-subarguments.go bar
subcommand 'bar'
  level: 0
  tail: []
$ go run i-flag-subarguments.go bar level
subcommand 'bar'
  level: 0
  tail: [level]
$ go run i-flag-subarguments.go bar level 0
subcommand 'bar'
  level: 0
  tail: [level 0]
$ go run i-flag-subarguments.go bar level 12312
subcommand 'bar'
  level: 0
  tail: [level 12312]
$ 
*/
