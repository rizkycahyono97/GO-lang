package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ===================== //
// JSON Object
// ===================== //

type Address struct {
	Street string
	Country string
	PostalCode string
}

type Customer struct {
	FirstName string
	MiddleName string
	LastName string
	Age int
	Married bool
	Hobbies []string	// array of string
	Address []Address   // array of object
}

type Customer2 struct {
	Name string `json:"name"`
	Languange string `json:"languange"`
	Id string `json:"id"`
	Bio string `json:"bio"`
	Version float64 `json:"version"`
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Rizky",
		MiddleName: "Cahyono",
		LastName: "Putra",
		Age: 22,
		Married: false,
	}

	// JSON Encoding
	bytes, _ := json.Marshal(customer)

	fmt.Print(string(bytes))
}