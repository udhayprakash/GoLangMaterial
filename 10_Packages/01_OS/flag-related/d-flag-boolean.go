package main

import (
    "flag"
    "fmt"
)

func main() {
    var flagvar bool
    flag.BoolVar(&flagvar, "flagvar", true, "Mandalorian Episode 4")
    fmt.Println("Boolean Value", flagvar)
}