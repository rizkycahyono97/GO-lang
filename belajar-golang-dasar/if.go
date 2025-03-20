package main

import "fmt"

func main() {
	name := "Rizky"

	if name == "Rizky" {
		fmt.Printf("Hello Rizky")
	} else if name == "Cahyono" {
		fmt.Println("Hello Cahyono")
	} else if name == "Putra" {
		fmt.Println("Hello Putra")
	} else { 
		fmt.Println("Nama kamu salah")
	}


	// =========== //
	// if short statement
	// =========== //

	if length := len(name); length >= 5 {
		fmt.Println("panjang nama sudah benar")
	} else {
		fmt.Println("panjang nama salah")
	}
}