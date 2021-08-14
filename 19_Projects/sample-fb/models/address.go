package models

// Address is used to store the values of address of account users
type Address struct {
	HouseNumber string `json:"housenumber,omitempty"`
	Street      string `json:"street,omitempty"`
	City        string `json:"city,omitempty"`
	Zipcode     string `json:"zipcode,omitempty"`
}
