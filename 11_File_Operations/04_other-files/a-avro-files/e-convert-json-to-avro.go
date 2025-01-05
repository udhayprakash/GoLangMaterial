package main

import (
	"fmt"

	"github.com/linkedin/goavro/v2"
)

func main() {
	// Define the Avro schema
	schema := `{
		"type": "record",
		"name": "User",
		"fields": [
			{"name": "name", "type": "string"},
			{"name": "age", "type": "int"},
			{"name": "email", "type": "string"}
		]
	}`

	// Create a new Avro codec
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		fmt.Println("Error creating codec:", err)
		return
	}

	// JSON data to convert
	jsonData := `{"name": "Alice", "age": 30, "email": "alice@example.com"}`

	// Convert JSON to Avro binary
	native, _, err := codec.NativeFromTextual([]byte(jsonData))
	if err != nil {
		fmt.Println("Error converting JSON to native:", err)
		return
	}

	binary, err := codec.BinaryFromNative(nil, native)
	if err != nil {
		fmt.Println("Error converting native to binary:", err)
		return
	}

	fmt.Printf("Avro binary data: %v\n", binary)
}
