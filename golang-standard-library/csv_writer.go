package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)	// akan ditulis di terminal / stdout

	_ = writer.Write([]string{"rizky", "cahyono", "putra"})		// _ =>  mengabaikan error
	_ = writer.Write([]string{"intan", "sukma", "mahsuni"})
	_ = writer.Write([]string{"adjie", "setya", "adi"})

	writer.Flush()
}