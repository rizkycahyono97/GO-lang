package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// ================== //
// pool
// ================== //
func TestPool(t *testing.T) {
	pool := sync.Pool{
		// default return for new (tidak wajib)
		New: func () interface{} {
			return "New"
		},
	}

	// membuat data di pool
	pool.Put("Rizky")
	pool.Put("Cahyono")
	pool.Put("Putra")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()	// ambil data pool dengan goroutine
			fmt.Println(data)
			time.Sleep(1 * time.Second)	// ada delay buat uji coba
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}