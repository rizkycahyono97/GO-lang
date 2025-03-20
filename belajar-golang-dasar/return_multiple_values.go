package main

import "fmt"

func getFullName() (string, string) {	// harus menuliskan kedua tipedata yang mau direturn
	return "eko", "Khaneddy" 
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	namaPertama, _ := getFullName()	// jika tidak ingin menggunakan salah satu nilai return / ignoring
	fmt.Println(namaPertama)

}