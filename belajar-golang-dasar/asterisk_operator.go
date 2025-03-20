package main

import "fmt"

type Address struct {
	City, Province, Region string
}

func main() {

	var address1  = Address{"Lamongan", "Jawa Timur", "Indonessia"}
	var address2  = &address1	

	address2.City = "Tuban"

	fmt.Println(address1)
	fmt.Println(address2)


	// address2 = &Address{"Boyolali", "Jawa Tengah", "Indonesia"} // hanya akan merubah di address2
	*address2 = Address{"Boyolali", "Jawa Tengah", "Indonesia"} // merubah semuanya di address1 dan juga di address2

	fmt.Println(address1)
	fmt.Println(address2)	
}