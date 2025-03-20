package main

import "fmt"

type Address struct {
	City, Province, Region string
}

func main() {

	// =================== //
	// copy value (boros memory)
	// =================== //

	// address1 := Address{"Lamongan", "Jawa Timur", "Indonessia"}
	// address2 := address1	// copy value from address1

	// address2.City = "Tuban"

	// fmt.Println(address1)
	// fmt.Println(address2)


	// =================== //
	// pointer / reference
	// =================== //

	address1 := Address{"Lamongan", "Jawa Timur", "Indonessia"}
	address2 := &address1	// pointer to address1
	address3 := new(Address)	// bisa juga menggunakan new -> pointer to Address

	address2.City = "Tuban"

	fmt.Println(address1)	// ikut berubah
	fmt.Println(address2)	// berubah juga menjadi Tuban
	fmt.Println(address3)	// berubah juga menjadi Tuban

}