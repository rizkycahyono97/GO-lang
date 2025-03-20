package main

import "fmt"

type Man struct {
	name string
}

func (man *Man) Married() {
	man.name = "Mr. " + man.name
}

func main() {
	rizky := Man{"Rizky"}	// mengisi nilai struct
	rizky.Married()

	fmt.Println(rizky.name)
}