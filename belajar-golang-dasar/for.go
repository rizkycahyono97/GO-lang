package main

import "fmt"

func main() {

	// bentuk pertama
	counter := 1;

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	fmt.Println("selesai")

	// bentuk kedua
	for counter2 := 1; counter2 <= 5; counter2++ {
		fmt.Println("perulangan ke ", counter2)
	}

	fmt.Println("selesai")

	// 1. contoh tipe data collection
	names := []string {"Rizky", "Cahyono", "Putra"}
	for i:=0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// 2. bentuk kedua (recomended)
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	// 3. jika tidak memakai index menggunakan _
	for _, name := range names {
		fmt.Println(name)
	}
}