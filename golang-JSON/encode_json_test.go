package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ===================== //
// JSON Encoding
// ===================== //
func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestLogJson(t *testing.T) {
	logJson("Hello World")
	logJson(123)
	logJson(123.456)
	logJson(true)	// boolean
	logJson([]string{"Hello", "World"})	// array
}