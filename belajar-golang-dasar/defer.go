package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {

	// defer = function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function dijalankan
	// misalkan error kode program diatasnya, maka defer tetap dijalankan
	defer logging()

	fmt.Println("Run Application")
}

func main() {
	runApplication()
}