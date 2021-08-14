package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Book struct {
	Name        xml.Name `xml:"book"`
	Id          string   `xml:"id,attr"`
	Title       string   `xml:"title"`
	Genre       string   `xml:"genre"`
	Price       float32  `xml:"price"`
	Date        string   `xml:"publish_date"`
	Description string   `xml:"description"`
}

func (b *Book) GetId() string {
	return b.Id
}

func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) GetGenre() string {
	return b.Genre
}

func (b *Book) GetPrice() float32 {
	return b.Price
}

func (b *Book) GetDate() string {
	return b.Date
}

func (b *Book) GetDescription() string {
	return b.Description
}

func main() {
	file, err := os.Open("./books2.xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteData, _ := ioutil.ReadAll(file)

	var books Book
	xml.Unmarshal([]byte(byteData), &books)
	fmt.Println(books)

	for i := range books {
		fmt.Println(i)
	}
	// for i := 0; i < len(books.Book); i++ {
	// 	fmt.Println("\nUser Type: " + books.Book[i].Type)
	// 	fmt.Println("User Name: " + books.Book[i].Name)
	// 	fmt.Println("Facebook Url: " + books.Book[i].Social.Facebook)
	// }

}
