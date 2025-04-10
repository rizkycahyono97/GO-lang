package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ===================== //
// JSON Decoding
// ===================== //

func TestDecodeJson(t *testing.T) {
	jsonString := `{"FirstName":"Rizky","MiddleName":"Cahyono","LastName":"Putra","Age":22,"Married":false}`
	jsonBytes := []byte(jsonString)

	// tempat menyimpan hasil decode
	customer := &Customer{}

	// JSON Decoding
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	// Print the result
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Age)
}