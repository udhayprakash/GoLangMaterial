package main

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

var standardPackages = make(map[string]struct{})

func init() {
	pkgs, err := packages.Load(nil, "std")
	if err != nil {
		panic(err)
	}

	for _, p := range pkgs {
		standardPackages[p.PkgPath] = struct{}{}
	}
}

func isStandardPackage(pkg string) bool {
	_, ok := standardPackages[pkg]
	return ok
}

func main() {
	fmt.Println(isStandardPackage("fmt"))  // true
	fmt.Println(isStandardPackage("nope")) // false
}

/*
3rd party packages import
	1. Relative path import

		“./model” // load package in the same directory, I don’t recommend this way.

	2. Absolute path import

		“shorturl/model” // load package in path “$GOPATH/pkg/shorturl/-model”
*/
