package main

import "fmt"

func getGoodBye(name string) string {
	return "Hello " + name
}

func main() {

	// function disimpan di variable / function as value
	contoh := getGoodBye
	misal := getGoodBye
	
	fmt.Println(contoh("Rizky"))
	fmt.Println(misal("Yanto"))
}