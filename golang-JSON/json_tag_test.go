package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ====================== //
// JSON Tag
// ====================== //
// JSON Tag = untuk mengatur nama field pada JSON
// yang sebelumnya dari CamelCase menjadi snake_case(standard JSON)
type Product struct{
	Name string `json:"name"`
	Price int `json:"price"`
	Category string `json:"category"`
}

// encode
func TestJSONTag(t *testing.T) {
	product := Product{
		Name: "Laptop",
		Price: 10000000,
		Category: "Electronics",
	}

	// JSON Encoding
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

// decode
func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"name":"Laptop","price":10000000,"category":"Electronics"}`
	jsonBytes := []byte(jsonString)

	// tempat menyimpan hasil decode
	product := &Product{}

	// JSON Decoding
	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}