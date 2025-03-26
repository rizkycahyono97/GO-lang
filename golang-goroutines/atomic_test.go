package golanggoroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// =============== //
// Atomic = kebalikan race condition
// =============== //
func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				// x = x + 1 
				atomic.AddInt64(&x, 1)	// sama seperti atas tetapi lebih baik, tidak race condition
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter ", x)
}