package main

import (
	"container/ring" // struktur data circular list
	"fmt"
	"strconv"
)

func main() {

	var data *ring.Ring = ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	}) 


	// =============== //
	// cara manual
	// =============== //
	// data.Value = "Data 1"

	// data = data.Next()
	// data.Value = "Data 2"

	// data = data.Next()
	// data.Value = "Data 3"

	// data = data.Next()
	// data.Value = "Data 4"

	// data = data.Next()
	// data.Value = "Data 5"

	data.Do(func(value any) {
		fmt.Println(value)
	}) 

}