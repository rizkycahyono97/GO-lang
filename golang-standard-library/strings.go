package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Rizky Cahyono", "Cahyo"))
	fmt.Println(strings.Split("Rizky Cahyono", ""))		// menjadikan split
	fmt.Println(strings.ToLower("RIZKY CAHYONO"))
	fmt.Println(strings.ToUpper("rizky cahyono"))
	fmt.Println(strings.Trim("            Rizky Cahyono           ", " "))	// menghilangkan spasi
	fmt.Println(strings.ReplaceAll("Rizky Cahyono Rizky Cahyono", "Rizky", "Budi"))		// mereplace string 
}