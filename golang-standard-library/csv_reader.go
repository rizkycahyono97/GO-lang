package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "rizky,cahyono,putra\n" +
					"intan,sukma,mahsuni\n" +
					"nugraha,abuyis,nahsa"

	reader := csv.NewReader(strings.NewReader(csvString))

	// loop untuk membaca semua csv
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}