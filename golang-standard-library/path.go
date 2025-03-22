package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))		// extension file
	fmt.Println(path.Join("hello", "world", "main.go"))		// menggabungkan menjadi path

}