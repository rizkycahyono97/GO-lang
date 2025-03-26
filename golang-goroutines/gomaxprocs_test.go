package golanggoroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// ================ //
// gomacprocs = untuk mengecek thread dll
// ================ //
func TestGomacprocs(t *testing.T)  {
	// test total goroutine
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(5 * time.Second)
			group.Done()
		}()
	}


	// total CPU
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU : ", totalCpu)

	// total thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Threads: ", totalThread)

	// total goroutine
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine : ", totalGoroutine)

	group.Wait()
}

func TestChangeThread(t *testing.T)  {
	// test total goroutine
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(5 * time.Second)
			group.Done()
		}()
	}

	// total CPU
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU : ", totalCpu)

	// total thread
	runtime.GOMAXPROCS(10)	// kita rubah thread menjadi 10
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Threads: ", totalThread)

	// total goroutine
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine : ", totalGoroutine)

	group.Wait()
}