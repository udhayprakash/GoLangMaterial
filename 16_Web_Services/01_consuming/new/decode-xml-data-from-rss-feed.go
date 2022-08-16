package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

func main() {

	response, err := http.Get("http://www.thestar.com.my/RSS/Metro/Community/")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer response.Body.Close()

	XMLdata, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rss := new(Rss)

	buffer := bytes.NewBuffer(XMLdata)

	decoded := xml.NewDecoder(buffer)

	err = decoded.Decode(rss)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Title : %s\n", rss.Channel.Title)
	fmt.Printf("Description : %s\n", rss.Channel.Description)
	fmt.Printf("Link : %s\n", rss.Channel.Link)

	total := len(rss.Channel.Items)

	fmt.Printf("Total items : %v\n", total)

	for i := 0; i < total; i++ {
		fmt.Printf("[%d] item title : %s\n", i, rss.Channel.Items[i].Title)
		fmt.Printf("[%d] item description : %s\n", i, rss.Channel.Items[i].Description)
		fmt.Printf("[%d] item link : %s\n\n", i, rss.Channel.Items[i].Link)
	}

}
