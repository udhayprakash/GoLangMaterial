package main

 import (
         "fmt"
         "sort"
 )

 var strSlice sort.StringSlice = []string{"apple", "durian", "kiwi", "banana"}

 func main() {

         fmt.Println("Original : ", strSlice[:])

         strSlice.Sort()

         fmt.Println("Sort : ", strSlice[:])

         sort.Sort(sort.Reverse(strSlice[:]))    

         fmt.Println("Reverse : ", strSlice[:])

 }