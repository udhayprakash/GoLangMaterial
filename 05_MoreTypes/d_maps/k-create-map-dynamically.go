package main

import (
	"fmt"
	"reflect"
)

func main() {
	key := 1
	value := "abc"

	mapType := reflect.MapOf(reflect.TypeOf(key), reflect.TypeOf(value))

	mapValue := reflect.MakeMap(mapType)
	mapValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
	mapValue.SetMapIndex(reflect.ValueOf(2), reflect.ValueOf("def"))
	mapValue.SetMapIndex(reflect.ValueOf(3), reflect.ValueOf("gh"))

	keys := mapValue.MapKeys()
	for _, k := range keys {
		c_key := k.Convert(mapValue.Type().Key())
		c_value := mapValue.MapIndex(c_key)
		fmt.Println("key :", c_key, " value:", c_value)
	}
	fmt.Printf("%T:  %v\n", mapValue.Interface(), mapValue.Interface())
	fmt.Println("mapValue =", mapValue)


}
