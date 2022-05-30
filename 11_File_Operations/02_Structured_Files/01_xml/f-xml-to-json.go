package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Data struct {
	XMLName    xml.Name `xml:"data" json:"-"`
	PersonList []Person `xml:"person" json:"people"`
}

type Person struct {
	XMLName   xml.Name `xml:"person" json:"-"`
	Firstname string   `xml:"firstname" json:"firstname"`
	Lastname  string   `xml:"lastname" json:"lastname"`
	Address   *Address `xml:"address" json:"address,omitempty"`
}

type Address struct {
	City  string `xml:"city" json:"city,omitempty"`
	State string `xml:"state" json:"state,omitempty"`
}

func main() {
	rawXmlData := "<data><person><firstname>Nic</firstname><lastname>Raboy</lastname><address><city>San Francisco</city><state>CA</state></address></person><person><firstname>Maria</firstname><lastname>Raboy</lastname></person></data>"
	var data Data
	xml.Unmarshal([]byte(rawXmlData), &data)
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
}
