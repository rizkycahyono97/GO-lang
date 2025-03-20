package helper

import "fmt"

var version = "1.0"		// tidak bisa diakses di package selain helper
var Application = "Go"

// harus membuatnya dg huruf besar didepan, supaya bisa diakses di package lain
func SayHello(name string) string {
	return "Hello " + name
}

// tidak bisa diakses di package selain helper
func sayGoodBye(name string) string {
	return "Good Bye " + name
}

// bisa diakses jika sesama package
func Contoh() {
	fmt.Println(version)
}