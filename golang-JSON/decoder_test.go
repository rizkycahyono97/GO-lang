package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// 	===================== //
// JSON Encoding 
// 	===================== //
// untuk membaca file JSON 

func TestJSONDecoder(t *testing.T) {

	reader, err  := os.Open("Customer2.json")
	if err != nil {
		panic(err)
	}
	defer reader.Close()	// Close the file after reading

	decoder := json.NewDecoder(reader)  // Read JSON File

	var customer2 []Customer2  // Decode JSON to struct
	err = decoder.Decode(&customer2)	// seperti Func UnMarshal
	if err != nil {
		panic(err)
	}

	// Print the result
	fmt.Println(customer2)
}