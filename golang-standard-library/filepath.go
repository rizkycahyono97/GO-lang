package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	// ============= //
	// khusus windows biasanya
	// ============= //

	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go"))	// kalau diwindows dari folder C:
	fmt.Println(filepath.IsLocal("hello/world.go"))
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}