package main

import (
	"flag"
	"fmt"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "dev", "Project Environment")

	flag.Parse()

	fmt.Println("Environment =", env)
}

/*
OUTPUT:
-------
	~go run c-string_flags.go
	Environment = dev

	~go run c-string_flags.go -env=dev
	Environment = dev~go run c-string_flags.go -env=prod
	Environment = prod

	~go run c-string_flags.go -env=prod uat stage
	Environment = prod

	~go run c-string_flags.go -h
	Usage of C:...\c-string_flags.exe:
	-env string
			Project Environment (default "dev")
*/
