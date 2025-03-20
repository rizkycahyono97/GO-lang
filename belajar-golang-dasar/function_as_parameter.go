package main

import "fmt"

// jika function diparameter terlalu panjang, kita bisa memakai alias / type declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {		// didalam parameter function merupakan function juga	
	filterName := filter(name)	// deklarasi function di parameter

	fmt.Println("Hello", filterName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}
}

func main() {
	// cara memanggilnya
	sayHelloWithFilter("Rizky", spamFilter, )

	filtered := spamFilter
	sayHelloWithFilter("Anjing", filtered)
}