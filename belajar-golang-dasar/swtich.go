package main

import "fmt"

func main() {
	name := "Rizky"

	switch name {
	case "Rizky":
		fmt.Println("Hello Rizky")

	case "Eko":	
		fmt.Println("Hello Eko")

	default:
		fmt.Println("nama salah")
	}

	// ============== //
	// swith short statement
	// ============== //
	switch lenght := len(name); lenght > 7 {
	case true:
		fmt.Println("Hello Rizky")

	case false:
		fmt.Println("nama salah")
	}


	// ============== //
	// swith tanpa kondisi
	// ============== //
	lenght := len(name)
	switch {
	case lenght > 10:
		fmt.Println("nama terlalu panjang")

	case lenght > 5:
		fmt.Println("nama lumayan pas")

	default:
		fmt.Println("nama pas")
	}
}