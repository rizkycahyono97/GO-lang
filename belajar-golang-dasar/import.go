package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Rizky")

	fmt.Println(result)

	// fmt.Println(helper.version)	// tidak bisa diakses
	// fmt.Println(helper.sayGoodBye)	// tidak bisa diakses
}

