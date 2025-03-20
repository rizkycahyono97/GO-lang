package main

import (
	"fmt"
)

type validationError struct {
	Message string
}

// custom error
func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

// custom error
func (n *notFoundError) Error() string {
	return n.Message 
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Id tidak boleh kosong"}
	} else if id != "Rizky" {
		return &notFoundError{"nama tidak valid"}
	}

	return nil 	// return nil jika tidak ada error
}

func main() {
	err := SaveData("", nil)

	if err != nil {
		// terjadi error

		// ika memakai if else
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation Error, ", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not Found Error, ", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknown Error")
		// }

		// menggunakan switch case
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error, ", finalError.Error())
		case *notFoundError:
			fmt.Println("Not Found Error, ", finalError.Error())
		default:	
			fmt.Println("Unknown Error")
		}
		
	} else {
		fmt.Println("Sukses")
	}
}