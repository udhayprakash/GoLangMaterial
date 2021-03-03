package main

type MyType struct {
	Value string
}

func (mt MyType) SetValue(value string) {
	mt.Value = value
}

func (mt *MyType) SetValuePtr(value string) {
	mt.Value = value
}

func main() {
	mt := MyType{Value: "foo"}
	mt.Value = "bar"
	mt.SetValue("baz")
	println("mt.Value:", mt.Value) // 'bar

	mt.Value = "bar"
	mt.SetValuePtr("baz")
	println("mt.Value:", mt.Value) // baz
}
