package main

import "fmt"

func endApp() {
	fmt.Println("End App")

	// recover = akan tetap menjalankan kodeprogram ketika panic / counter panic
	// yang benar kita taruh recover di defernya / di defer saat ini adalah defer endApp()
	message := recover()
	fmt.Println("Terjadi error", message)
}

func runApp(error bool) {
	defer endApp()	
	
	// panic adalah akan mengehentikan kode program
	if error {
		panic("Ups Error")
	}
}

func main() {
	runApp(true)	// karena di func runApp kita memakai defer maka tetap akan dijalankan meskipun memakai panic()
	fmt.Println("Rizky Cahyono Putra")
}

// defer, panic, recover seperti try and catch