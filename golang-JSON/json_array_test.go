package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ====================== //
// JSON Array
// ====================== //

// encode
func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "Rizky",
		MiddleName: "Cahyono",
		LastName: "Putra",
		Age: 22,
		Married: false,
		Hobbies: []string{"Coding", "Reading", "Gaming"},
	}

	// JSON Encoding
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}


// Decode
func TestDecodeJSONArray(t *testing.T) {
	jsonString := `{"FirstName":"Rizky","MiddleName":"Cahyono","LastName":"Putra","Age":22,"Married":false,"Hobbies":["Coding","Reading","Gaming"]}`
	jsonBytes := []byte(jsonString)

	// tempat menyimpan hasil decode
	customer := &Customer{}

	// JSON Decoding
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}




// ======================= //
// JSON Array Complex / Array Object
// ======================= //
// Array Complex = object JSON didalam Array
func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Rizky",
		Address: []Address{			// array of object
			{
				Street: "Jl. Raya No. 1",
				Country: "Indonesia",
				PostalCode: "12345",
			},
			{
				Street: "Jl. Raya No. 2",
				Country: "Indonesia",
				PostalCode: "67890",
			},
		},
	}

	// JSON Encoding
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

// decode
func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Rizky","MiddleName":"Cahyono","LastName":"Putra","Age":22,"Married":false,"Address":[{"Street":"Jl. Raya No. 1","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jl. Raya No. 2","Country":"Indonesia","PostalCode":"67890"}]}`
	jsonBytes := []byte(jsonString)

	// tempat menyimpan hasil decode
	customer := &Customer{}

	// JSON Decoding
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Address)
}


// ======================= //
// Array to JSON encode & decode

// decode form array biasa to object/json
func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jl. Raya No. 1","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jl. Raya No. 2","Country":"Indonesia","PostalCode":"67890"}]`
	jsonBytes := []byte(jsonString)

	// tempat menyimpan hasil decode
	addresses := &[]Address{}

	// JSON Decoding
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
	fmt.Println((*addresses)[0].Street)
}


// encode
func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{			// array of object
		{
			Street: "Jl. Raya No. 1",
			Country: "Indonesia",
			PostalCode: "12345",
		},
		{
			Street: "Jl. Raya No. 2",
			Country: "Indonesia",
			PostalCode: "67890",
		},
	}

	// JSON Encoding
	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}