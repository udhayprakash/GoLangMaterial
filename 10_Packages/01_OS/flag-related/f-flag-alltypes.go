package main

import (
    "flag"
    "fmt"
)

func main() {
    var flagvar int
    var flagvar2 string
    var flagvar3 bool

    flag.IntVar(&flagvar, "flagvar", 111921, "Mandalorian Episode 4")
    flag.StringVar(&flagvar2, "flagvar2", "Gina Carano", "Mandalorian Episode")
    flag.BoolVar(&flagvar3, "flagvar3", true, "Mandalorian")

    fmt.Println("Integer Value:", flagvar)
    fmt.Println("String Value :", flagvar2)
    fmt.Println("Boolean Value:", flagvar3)
}