package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name": name,
		}
	}
}

func main() {
	data := newMap("")

	if data == nil {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Print(data["name"])
	}

	// fmt.Println(data[""])	// akan mengembalikan nilai nil
	// fmt.Println(data["name"])
}