package main

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

func main() {

	type Company struct {
		XMLName xml.Name `xml:"company"`
	}

	c := Company{}
	cType := reflect.TypeOf(c)
	field := cType.Field(0)

	xmlTag := field.Tag.Get("xml")
	fmt.Println(xmlTag)

	type S struct {
		F string `species:"gopher" color:"pink"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field = st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

}
