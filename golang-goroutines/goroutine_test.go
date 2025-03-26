package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutines(t *testing.T) {
	go RunHelloWorld()	// go(goroutine) berjalan sebagai async
	fmt.Println("Test")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(num int) {
	fmt.Println("Display ", num)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)		// jika device menggunakan multithread maka angka tidak akan urut
	}

	time.Sleep(5 * time.Second)
}