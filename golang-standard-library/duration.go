package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 = duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("Duration : %d\n", duration3)
}