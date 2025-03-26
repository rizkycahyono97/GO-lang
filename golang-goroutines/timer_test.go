package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// ================== //
// timer
// ================== //

// bentuk pertama
func TestTime(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)		// delay 5 detik
	fmt.Println(time.Now())

	time := <- timer.C	// channel
	fmt.Println(time)
}

// bentuk kedua
func TestAfterTime(t *testing.T) {
	channel := time.After(5 * time.Second)		// delay 5 detik
	fmt.Println(time.Now())

	time := <- channel// channel
	fmt.Println(time)
}


// bentuk ketiga (after func)
func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	
	// after func memakai goroutine
	time.AfterFunc(5 * time.Second, func ()  {
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()
}