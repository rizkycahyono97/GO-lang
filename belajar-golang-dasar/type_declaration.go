package main

import "fmt"

func main() {
	
	type noKtp string	// alias bahwa noKtp adalah tipe data string

	var ktpEko noKtp = "123534456"

	var contoh string = "2222222222222"
	var contohKtp noKtp = noKtp(contoh)		// assign

	fmt.Println(ktpEko)
	fmt.Println(contohKtp)
}