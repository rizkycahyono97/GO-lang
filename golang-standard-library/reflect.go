package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name	string `required:"true" max:"10"`
	Email	string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
}

// function untuk mengetahui / reflect type dan field dari struct
func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)	// ambil valuenya di paramater
	fmt.Println("Type Name: ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var StructField reflect.StructField = valueType.Field(i)
		fmt.Println(StructField.Name, "With Type", StructField.Type)
		fmt.Println(StructField.Tag.Get("required"))	// untuk mengetahui tagnya
		fmt.Println(StructField.Tag.Get("max"))
	}
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)	// ambil data
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {	// jika required == true
			data := reflect.ValueOf(value).Field(i).Interface()	// print data
			result = data != ""		// jika tag kosong
			if result == false {	// akan tetap looping sampai result == false
				return	result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"rizky"})
	readField(Person{"rizky", "", ""})

	person := Person {
		Name: "Rizky",
		Email: "ada",
		Address: "",
	}

	fmt.Println(isValid(person))
}