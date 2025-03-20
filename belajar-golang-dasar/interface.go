package main

import "fmt"

// interface = sekumpulan method tanpa implementasi, tetapi harus ada di func lain secara lengkap
type hasName interface {
	getName() string
}

// mengambil method dari interface, dan harus di ambil semua
func sayHello(value hasName) {
	fmt.Println("Hello", value.getName())
}

// struct
type Person struct {
	Name string
}

type Hewan struct {
	Name string
}

// function untuk menerima interface
func (hewan Hewan) getName() string {
	return hewan.Name
}

func (person Person) getName() string { 
	return person.Name
}

func main() {
	person := Person{Name: "Rizky"}
	sayHello(person)

	hewan := Hewan{Name: "Kucing"}
	sayHello(hewan)
}