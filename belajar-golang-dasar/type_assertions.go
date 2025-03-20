package main

import "fmt"

// type_Assertion adalah konversi tipe data

func random() any {
	return "ok"
} 

func main() {
	var result any = random() // akan mengembalikan nilai any

	// ============== //
	// bentuk manual
	// ============== //
	var resultString string = result.(string)  // konversi dari any ke string
	// var resultInt int = result.(int)	// akan menghasilkan panic, karena data aslinya string
	fmt.Println(resultString)


	// ============== //
	// bentuk type_assertion_switch
	// ============== //
	switch value := result.(type) {		// cari typenya apa
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("default", value)
	}
}