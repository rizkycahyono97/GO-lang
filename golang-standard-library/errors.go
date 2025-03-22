package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation Error")
	NotFoundError = errors.New("Not Found Error")
)

func getId(id string) error {
	if id == "" {
		return ValidationError
	} else if id != "Rizky" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := getId("")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error") 
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("Unknown Error")
		}
	} else {
		fmt.Println("Success")
	}
}