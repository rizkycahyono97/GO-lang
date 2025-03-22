package main

import (
	"fmt"
	"slices"
)

func main() {
	name := []string{"rizky", "eko", "rara", "aura"}
	values := []int{10, 20, 50, 45, 90}

	fmt.Println(slices.Min(name))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(name))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(name, "rara"))
	fmt.Println(slices.Index(name, "aura"))
	fmt.Println(slices.Index(name, "John"))


}