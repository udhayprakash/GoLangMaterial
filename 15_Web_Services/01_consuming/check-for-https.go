package main

 import (
 	"fmt"
 	"net/http"
 	"strings"
 )

 func main() {

 	// test websites
 	//originalURL := "//socketloop.com"
 	//originalURL := "http://geocities.com"  -- no https
 	originalURL := "http://cowner.net"  // -- no https


 	resp, err := http.Get(originalURL)

 	if err != nil {
 		fmt.Println(err)
 	}

 	// if there is any re-direction happening behind the scene
 	// the finalURL will be different
 	// in this case, there will be a re-direction to https (SSL) version

 	finalURL := resp.Request.URL.String()

 	fmt.Println("Original URL is : ", originalURL)
 	fmt.Println("Final URL is : ", finalURL)

 	// Check if served with https 
 	fmt.Println("Is HTTPS ? : ", strings.HasPrefix(finalURL,"https"))

 }