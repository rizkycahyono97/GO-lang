package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eKo"))
	fmt.Println(regex.MatchString("e1o"))
	fmt.Println(regex.MatchString("eyo"))


	fmt.Println(regex.FindAllString("eko eKo ego e1o", 10))
}