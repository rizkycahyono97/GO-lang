package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {

		// jika i bisa di modulus 2, maka lanjutkan jangan di print
		if i%2 == 0 {
			continue
		}

		fmt.Println("angka ganjil ke", i)
	}
}