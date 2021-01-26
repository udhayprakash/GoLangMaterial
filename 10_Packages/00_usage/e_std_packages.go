package main

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

func main() {
	pkgs, err := packages.Load(nil, "std")
	if err != nil {
		panic(err)
	}
	for _, p := range pkgs {
		fmt.Println(p)
	}
	// Output: [archive/tar archive/zip bufio bytes compress/bzip2 ... ]
}
