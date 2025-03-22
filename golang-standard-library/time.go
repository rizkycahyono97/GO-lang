package main

import (
	"fmt"
	"time"
)

func main() {

	// time now
	var now time.Time = time.Now()
	fmt.Println(now)

	// UTC
	var utc time.Time = time.Date(2009, time.April, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	// 
	formatter := "2006-01-02 15:04:05"	// format standar GO 

	value := "2005-05-06 10:10:10"	// misalkan inputan user

	valueTime, err := time.Parse(formatter, value)	// parsing
	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
}