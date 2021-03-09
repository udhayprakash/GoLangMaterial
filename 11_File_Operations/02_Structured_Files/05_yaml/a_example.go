package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

var data = []byte(`- key1: item 1
  key2: item 2
- key1: another item 1
  key2: another item 2`)

type ListItem struct {
	Line int

	ListItemData
}

type ListItemData struct {
	Key1 string
	Key2 string
}

func (li *ListItem) UnmarshalYAML(value *yaml.Node) error {
	err := value.Decode(&li.ListItemData)
	if err != nil {
		return err
	}

	li.Line = value.Line

	return nil
}

func main() {
	var list []ListItem

	err := yaml.Unmarshal(data, &list)
	if err != nil {
		panic(err)
	}

	for _, i := range list {
		fmt.Printf("%#v\n", i)
	}
}
