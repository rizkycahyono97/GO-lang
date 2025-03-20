package main

import "fmt"

// vararg function
func sumAll(numbers ...int) int {	// ...int menunjukan parameter diisi dengan tipe data slice
	total := 0

	// untuk mengihutung total
	for _, number := range numbers {
		total += number
	} 

	return total
}

func main() {
	fmt.Println(sumAll(10, 20, 30, 40))
	fmt.Println(sumAll(2, 3))

	// aslinya, tidak usah membikin slice lagi
	// fmt.Println(sumAll([]int{1, 2, 3}))

	// jika terlanjur punya slice
	numbers := []int{10, 20, 30}
	fmt.Println(sumAll(numbers...))	// harus menambahkan ... diakhir 
}