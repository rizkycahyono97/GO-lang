package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

// =============== //
// race conditional = karena gotoutine concurency maka akan terjadi racing, variabel tidak akan di eksekusi semua
// =============== //
func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1	// x akan ditambahkan 1 setiap perulangan / manipulasi variable
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter ", x)
}