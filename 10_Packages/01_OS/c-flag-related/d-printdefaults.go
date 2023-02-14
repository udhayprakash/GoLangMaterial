package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "", "Project Environment")

	flag.Parse()

	if len(env) == 0 {
		// This condition will work only when default value is empty string
		fmt.Println("Usage: c-string_flags.go -name")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println("Environment =", env)
}

/*
OUTPUT:
-------

	~go run d-printdefaults.go -env=prod
	Environment = prod

	~go run d-printdefaults.go -h
	Usage of C:...\d-printdefaults.exe:  -env string
			Project Environment

	~go run d-printdefaults.go
	Usage: c-string_flags.go -name
	-env string
			Project Environment
	exit status 1


*/
