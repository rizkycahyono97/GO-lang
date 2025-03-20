package main

import "fmt"

func main() {
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)	// hasilnya akan -32767 karena batas int16


	// ============================ //
	var name = "Rizky Cahyono"
	var e = name[0]
	var eString = string(e)
	
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}