package main

import "fmt"

func getHello(name string) string {	// harus menuliskan tipe data untuk return
	hello := "Hello " + name

	return hello
}

func main() {

	// cara memanggil
	result := getHello("Agus")
	fmt.Println(result)

	fmt.Println(getHello("Cahyono"))
}