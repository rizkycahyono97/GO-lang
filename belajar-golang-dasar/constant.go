package main

import "fmt"

func main() {

	// constant tidak bisa dirubah, dan tidakwajib assign
	const firstName string = "Rizky"
	const lastName = "Cahyono"
	const (	// multiple const
		fullName = "Intan Sukma"
		middleName = "Cahyono"
	)

	// error
	// firstName = "Intan"
	// lastName = "Intan"

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(fullName)
	fmt.Println(middleName)
}