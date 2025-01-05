package main

import (
	"fmt"
	"os"

	"github.com/linkedin/goavro/v2"
)

/*
Avro file format 	: https://avro.apache.org/
Installation		: github.com/linkedin/goavro/v2

.avro	Binary avro data files
		conatins serialized data in binary format
		includes schema in header (when using object container file format)

.avsc	Avro schema files(json format)
		Defines structure of avro data

.avdl	Avro IDL (Interface Definition Language) files
		Defines avro protocols in human-readable syntax
		Used for RPC (remote procedure call) definitions

.avpr	avro protocol files (json format)
		Defines Avro protocols in json format
		simalar to .avsc, but used in RPC definitions
*/

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

	// Create a new Avro file
	file, err := os.Create("users.avro")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create an Avro OCF (Object Container File) writer
	ocfWriter, err := goavro.NewOCFWriter(goavro.OCFConfig{
		W:      file,
		Codec:  codec,
		Schema: schema,
	})
	if err != nil {
		fmt.Println("Error creating OCF writer:", err)
		return
	}

	// Write data to the Avro file
	records := []map[string]interface{}{
		{"name": "Alice", "age": 30, "email": "alice@example.com"},
		{"name": "Bob", "age": 25, "email": "bob@example.com"},
	}
	if err := ocfWriter.Append(records); err != nil {
		fmt.Println("Error writing records:", err)
		return
	}

	fmt.Println("Data written to users.avro successfully!")
}
