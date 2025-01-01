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
------

$ go run b-string_flags.go
Environment = dev

$ go run b-string_flags.go -h
Usage of /tmp/go-build3701618246/b001/exe/b-string_flags:
  -env string
        Project Environment (default "dev")

$ go run b-string_flags.go -env=dev
Environment = dev

$ go run b-string_flags.go -env=prod
Environment = prod

$ go run b-string_flags.go -env=prod uat stage
Environment = prod

*/
