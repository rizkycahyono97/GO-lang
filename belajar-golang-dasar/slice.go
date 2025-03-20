package main

import (
	"fmt"
)

func main() {
	names := [...]string{
		"Rizky",
		"Cahyono",
		"Putra",
		"Intan",
		"Sukma",
		"Mahsuni",
	}

	// ================ //
	// penerapan slice
	// ================ //

	slice := names[4:6]	// index ke [4][5]
	fmt.Println(slice)

	slice2 := names[:3]	// dari index [0] sampai [2]
	fmt.Println(slice2)

	slice3 := names[3:]	// dari index [3] sampai terakhir
	fmt.Println(slice3)

	// model biasa
	var slice4 []string = names[:]	// untuk semua index 
	fmt.Println(slice4)

	// ================ //
	// function-funtion di slice
	// ================ //

	days := [...]string { 
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"ahad",
	}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	// assign
	daySlice1[0] = "sabtu baru"	// akan assign ke index[0] dari index slice bukan index array aslinya
	daySlice1[1] = "ahad baru"

	fmt.Println(days)

	// 1. append
	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"	// akan merubah di daySlice1[0] saja, bukan di array aslinya
	// tidak akan bertambah untuk array aslinya ketika di append, karena sifat dasar array yang tidak bisa ditambah
	// jadi seolah-olah daySlice2 membuat array baru dan ditambahkan "Libur Baru"
	fmt.Println(daySlice2)	
	fmt.Println(daySlice1)
	fmt.Println(days)

	// 2. make  
	// 2 -> panjang array, 5 -> kapasitas array, slice tidak akan membuat array baru sampai kapasitasnya mentok
	var newSlice []string = make([]string, 2, 5)	     
	newSlice[0] = "Rizky"
	newSlice[1] = "Cahyono"
	// newSlice[2] = "Putra" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))	// cek panjang
	fmt.Println(cap(newSlice))	// cek kapasitas

	newSlice2 := append(newSlice, "Putra")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))	// cek panjang
	fmt.Println(cap(newSlice2))	// cek kapasitas

	// coba logics yang lain
	newSlice2[0] = "Intan"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)


	// 3. copy
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)



	// =================== //
	// perbedaan array vs slice
	// =================== //

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}