package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Print("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Rizky", "Cahyono")
	sayHelloTo("Agus", "Cahyono")
}