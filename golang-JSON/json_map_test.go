package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ===================== //
// JSON Map
// ===================== //
// jika data yang kita terima tidak memiliki struktur yang tetap, kita menggunakan map

// decode
func TestJSONMap(t *testing.T) {
	jsonString := `{"name":"Laptop","price":10000000,"category":"Electronics"}`
	jsonBytes := []byte(jsonString)

	// tempat menyimpan hasil decode di tipe data map
	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["name"])
	fmt.Println(result["price"])
	fmt.Println(result["category"])
}

// encode
func TestJSONMapEncode(t *testing.T) {
	product := map[string]interface{} {
		"name": "Laptop",
		"price": 10000000,
		"category": "Electronics",
	}

	// JSON Encoding
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}