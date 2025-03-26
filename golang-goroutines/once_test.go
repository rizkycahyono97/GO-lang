package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

// ================ //
// onlyOnce
// ================ //
var counter = 0

func OnlyOne() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}			// once = hanya bisa mengakses func sekali saja
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			once.Do(OnlyOne)	// Do = melakukan once
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)		// hasilnya akan 1, kerena memakai onlyOnce
}
