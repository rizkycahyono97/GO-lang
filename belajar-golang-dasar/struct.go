package main

import "fmt"

// =============== //
// bentuk struct biasa
// =============== //

type Customer struct {
	Name, Address string
	Age int
}

// =============== //
// struct method
// =============== //
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello ", name, "My name is ", customer.Name)
}


func main() {

	// akses struct harus buat object dahulu
	var rizky Customer

	// deklarasi field
	rizky.Name = "Rizky"
	rizky.Address = "Lamongan"
	rizky.Age = 20

	fmt.Println(rizky)
	fmt.Println(rizky.Name)
	fmt.Println(rizky.Age)
	fmt.Println(rizky.Address)

	// =============== //
	// bentuk struct literals
	// =============== //
	// 1. pertama
	Joko := Customer{
		Name: "Joko",
		Address: "Indonesia",
		Age: 30,
	}
	fmt.Println(Joko)

	// 2/ kedua
	budi := Customer{"Budi", "Indonesia", 40}
	fmt.Println(budi)
}