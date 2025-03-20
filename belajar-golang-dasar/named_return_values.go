package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Rizky" // karena sudah dideklarasikan di atas maka tidak perlu pendeklarasian lagi
	middleName = "Cahyono"
	lastName = "Putra"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}