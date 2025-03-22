package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args		// untuk argument yang diinputkan user

	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Os Name: ", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
} 