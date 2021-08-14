package main

import (
    "fmt"
    "io/ioutil"
    "time"
)

func main() {
    err := ioutil.WriteFile("C:/Windows/test.txt", []byte("TESTING!"), 0644)
    if err != nil {
        fmt.Println(err.Error())
        time.Sleep(time.Second * 3)
    }
}