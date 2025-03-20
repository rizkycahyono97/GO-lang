package main

import "fmt"

type Address struct {
	City, Province, Region string
}

// pointer di func juga harus ada * di parameter yang dikirim
func changeCounteryToIndonesia(address *Address) {
	address.Region = "Indonesia"
}

func main() {
	var address Address = Address{}

	changeCounteryToIndonesia(&address)

	fmt.Println(address)
}