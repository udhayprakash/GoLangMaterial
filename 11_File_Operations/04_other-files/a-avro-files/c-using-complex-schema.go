package main

import (
	"fmt"
	"os"

	"github.com/linkedin/goavro/v2"
)

func main() {
	// Define a complex Avro schema
	schema := `{
		"type": "record",
		"name": "Person",
		"fields": [
			{"name": "name", "type": "string"},
			{"name": "age", "type": "int"},
			{"name": "address", "type": {
				"type": "record",
				"name": "Address",
				"fields": [
					{"name": "street", "type": "string"},
					{"name": "city", "type": "string"},
					{"name": "zipcode", "type": "string"}
				]
			}},
			{"name": "phoneNumbers", "type": {"type": "array", "items": "string"}}
		]
	}`

	// Create a new Avro codec
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		fmt.Println("Error creating codec:", err)
		return
	}

	// Create a new Avro file
	file, err := os.Create("people.avro")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create an Avro OCF writer
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
		{
			"name": "Alice",
			"age":  30,
			"address": map[string]interface{}{
				"street":  "123 Main St",
				"city":    "Springfield",
				"zipcode": "12345",
			},
			"phoneNumbers": []string{"123-456-7890", "987-654-3210"},
		},
	}
	if err := ocfWriter.Append(records); err != nil {
		fmt.Println("Error writing records:", err)
		return
	}

	fmt.Println("Data written to people.avro successfully!")
}
