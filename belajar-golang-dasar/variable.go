package main

import "fmt"

func main() {

	// cara-cara membuat variable,  jika tidak dipakai maka error
	var name string		// menggunakan tipedata
	var fullName = "Rizky Cahyono Putra"	// langsung menginisialisasikan
	firstName := "Intan"		// := seperti halnya deklarasi var, tetapi gunakan := di pertamakali membuat variable saja 
	var (	// multiple variable
		namaPertama = "Rizky"
		namaTengah = "Cahyono"
		namaTerakhir = "Putra"
	)

	name = "Rizky Cahyono"
	fmt.Println(name)

	name = "Rizky Putra"
	fmt.Println(name)

	fmt.Println(fullName)

	firstName = "Intan Sukma"
	fmt.Println(firstName)

	fmt.Println(namaPertama)
	fmt.Println(namaTengah)
	fmt.Println(namaTerakhir)

}