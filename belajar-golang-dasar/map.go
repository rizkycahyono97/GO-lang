package main

import "fmt"

func main() {

	// map[string]string => map[TypeKey]TypeValue

	// 1. bentuk pertama
	var person map[string]string = map[string]string {}
	person["name"] = "Rizky"
	person["address"] = "Lamongan"

	// 2. bentuk kedua (recomended)
	person2 := map[string]string {
		"name": "Rizky",
		"address": "Lamongan",
	}

	fmt.Println(person2["name"])
	fmt.Println(person2["address"])
	fmt.Println(person2)
	// fmt.Println(person2["salah"]) // jika mengakses yang tidak ada keynya


	// ============== //
	// function
	// ============== //
	book := make(map[string]string)	// membuat map

	book["title"] = "Bumi"
	book["author"] = "TereLiye"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}