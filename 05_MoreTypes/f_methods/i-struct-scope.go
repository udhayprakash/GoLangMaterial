package main

import "fmt"

type MyType struct {
	Value string
}

func (mt MyType) SetValue(value string) {
	// call  by value
	mt.Value = value
	fmt.Println("\nIn SetValue method, ", mt.Value)
}

func (mt *MyType) SetValuePtr(value string) {
	// call  by reference
	mt.Value = value
	fmt.Println("\nIn SetValuePtr method, ", mt.Value)
}

func main() {
	mt := MyType{Value: "foo"}

	// local changes
	mt.Value = "bar"
	println("mt.Value:", mt.Value)

	// Changes within method
	mt.SetValue("baz")
	println("mt.Value:", mt.Value) // 'bar

	// Changes within pointer method
	mt.SetValuePtr("baz")
	println("mt.Value:", mt.Value) // baz
}
