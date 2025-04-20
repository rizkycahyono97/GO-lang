package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	defer writer.Close()

	// JSON Encoding
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Rizky",
		MiddleName: "Cahyono",
		LastName: "Putra",
		Age: 22,
		Married: false,
	}

	// Encode JSON to file
	encoder.Encode(customer)	// Encode JSON to file
}
 