package main

import (
	"fmt"
	"strconv"
)

func main() {

	// string to boolean
	resultBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultBool)
	}

	// string to integer
	resultInt, err := strconv.Atoi("23")	
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	// int to binary
	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	// int to string
	stringInt := strconv.Itoa(999)
	fmt.Println(stringInt)
}