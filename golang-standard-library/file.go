package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// function create file
func createNewFile(name string, message string) error {

	// membuka file
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)	// name, flag(create|write only), permission
	if err != nil {
		return err
	}
	defer file.Close()	// jangan lupa file ditutup untuk menghindari memory leak
	file.WriteString(message)	// menulis 
	return nil
}

// function read and write file, tanpa menimpa yang sudah ada
func addToFile(name string, message string) error {

	// membuka file
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)	// name, flag(read write|appen), permission
	if err != nil {
		return err
	}
	defer file.Close()	// jangan lupa file ditutup untuk menghindari memory leak
	file.WriteString(message)	// menulis 
	return nil
}

// func read file
func readFile(name string) (string, error) {

	// membuka file
	file, err := os.OpenFile(name, os.O_RDONLY|os.O_APPEND, 0666)	// name, flag(read only|append)), permission
	if err != nil {
		return "", err
	}
	defer file.Close()

	// looping untuk membaca setiap barisnya
	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line)	+ "\n"
	}

	return message, err
}

func main() {
	// createNewFile("sample.log", "this is sample for write file")

	// result, _ := readFile("sample.log")
	// fmt.Println(result)

	addToFile("sample.log", "\nAdd new Message")
}