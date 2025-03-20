package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if (blackList(name)) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("You are Welcome", name)
	}
}

func main() {

	// 1. anonymous func model pertama
	blacklist := func (name string) bool {
		return name == "Anjing"
	}
	registerUser("Rizky", blacklist)


	// 2. anonymous func model kedua
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}