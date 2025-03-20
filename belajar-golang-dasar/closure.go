package main

import "fmt"

func main() {
	counter := 0

	// closure = kemampuan function berinteraksi dengan variable didekatnya
	increment := func ()  {
		fmt.Println("Increment")
		counter++
	}

	// closure
	increment()
	increment()
	increment()

	fmt.Println(counter)
}