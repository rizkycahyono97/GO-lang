package main

import "fmt"

func main() {

	// =============== //
	// bentuk pertama
	var names [3]string

	names[0] = "Rizky"
	names[1] = "Cahyono"
	names[2] = "Putra"

	names[1] = "Intan"
	fmt.Println(names[1])
	fmt.Println(names)


	// =================== //
	// bentuk kedua
	var names2 = [3]int {
		10,
		20,
		30,
	}

	fmt.Println(names2)
	fmt.Println(names2[1])


	// ================== //
	// bentuk ketiga
	var name3 = [...]int {		// [...] menandakan jumlah array yang tidak ditentukan
		10,
		20,
		30,
		40,
		50,
	}

	fmt.Println(len(name3))		// function len() untuk menghitung jumlah array
}