package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// ===================== //
// context
// ===================== //
func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}


// ==================== //
// context with value
// ==================== //
type contextKey string 

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()	// super parent
	
	contextB := context.WithValue(contextA, contextKey("b"), "B")	// (parent, key, value)
	contextC := context.WithValue(contextA, contextKey("c"), "C")	// (parent, key, value)

	// child dari contextB
	contextD := context.WithValue(contextB, contextKey("d"), "D")	
	contextE := context.WithValue(contextB, contextKey("e"), "E")

	// child dari contextC
	contextF := context.WithValue(contextC, contextKey("f"), "F")
	contextG := context.WithValue(contextC, contextKey("g"), "G")


	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)


	// get value
	fmt.Println(contextF.Value(contextKey("f")))
	fmt.Println(contextF.Value(contextKey("c")))	// akan dapat c, karena c parent dari f
	fmt.Println(contextF.Value(contextKey("b")))	// akan nil, karena b tidak ada di f maupun parentnya 
	fmt.Println(contextC.Value(contextKey("f")))	// akan nil, karena child hanya bisa mengakses parent atasnya bukan bawahnya	
}


// ===================== //
// context with cancel	= jika goroutine tidak mendapatkan data maka akan kita cancel, karena menyebabkan goroutine leak
// ===================== //

// goroutine leak
func CreateCounterLeak() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

// signal cancel (tidak leak)
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():		// akan menghentikan goroutine leak
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

// goroutine leak = terjadi ketika kita sudah menjalanakan goroutine, tetapi si goroutine masih berjalan, meskipun program sudah selesai
func TestContextWithCancelLeak(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())	// cek total goroutine

	destination := CreateCounterLeak()
	for n := range destination {
		fmt.Println("Total: ", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}

// tidak leak
func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
	for n := range destination {
		

		fmt.Println("Total: ", n)
		if n == 10 {
			break
		}
	}

	cancel()	// mengirim signal cancel ke context

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}



// ===================== //
// context with timeout = jika goroutine tidak mendapatkan data, maka akan tercance secara otomatis berdasarkan waktu
// ===================== //
func CreateCounterTimeout(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():		// akan menghentikan goroutine leak
				return
			default:
				destination <- counter
				time.Sleep(1 * time.Second)
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5 * time.Second)		// goroutine akan di cancel setelah 5 detik
	defer cancel()

	destination := CreateCounterTimeout(ctx)
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter: ", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}



// ================= //
// context with Deadline = seperti context with Timeout, tetapi kita mengasih waktu deadline untuk berhenti kapan
// ================= //
func CreateCounterDeadline(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():		// akan menghentikan goroutine leak
				return
			default:
				destination <- counter
				time.Sleep(1 * time.Second)
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second))		// goroutine akan di cancel setelah 5 detik
	defer cancel()

	destination := CreateCounterDeadline(ctx)
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter: ", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}