package main

type Options struct {
	Id      string `json:"id,omitempty"`
	Verbose bool   `json:"verbose,omitempty"`
	Level   int    `json:"level,omitempty"`
	Power   int    `json:"power,omitempty"`
}
/*
Cases to handle
---------------
- The JSON configuration can be missing some fields, and we'll want our Go struct
  to have default values for those.
- The JSON configuration can have extra fields which our struct doesn't have.
  Depending on the scenario, we may want to either ignore these or report an error.
*/


// ref: https://eli.thegreenplace.net/2020/optional-json-fields-in-go/