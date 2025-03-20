package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal" // blank identifier
	"fmt"
)

func main() {

	// akan otomatis menjalankan func init / constructor, meskipun kita memanggil GetDatabase
	fmt.Println(database.GetDatabase())
}