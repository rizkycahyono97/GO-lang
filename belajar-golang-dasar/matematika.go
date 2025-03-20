package main

import "fmt"

func main() {
	
	// untuk lambang operasi matematika seperti bahasa yang lain
	var a = 10
	var b = 15

	var c = a + b

	fmt.Println(c)

	// ===================== //
	var d = 10
	d += 20 // d = 20 + 10
	fmt.Println(d)

	var e = 15
	e /= 2 
	fmt.Println(e)


	// ======================= // 
	// unary operator
	var f = 1
	f++
	f++
	fmt.Println(f)

	var g = 1
	g--
	g--
	fmt.Println(g)

}