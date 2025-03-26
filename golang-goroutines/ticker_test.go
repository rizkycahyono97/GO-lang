package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

// =============== //
// ticker = untuk mengirimkan data/event secara berulang
// ============== //
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second) // akan mengirimkan data setiap 1 detik
	done := make(chan bool)

	// untuk stop
	go func() {
		time.Sleep(5 * time.Second) // setelah 5 detik berhenti
		ticker.Stop()               // untuk menghentikan
		close(done)
	}()

	for {
		select {
		case data := <- ticker.C:
			fmt.Println(data)
		case <- done:
			fmt.Println("Stopped")
			return
		}
	}
}


// =============== //
// tick = sama seperti ticker, tetapi hanya mengambil channelnya saja
// ============== //
func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second) // akan mengirimkan data setiap 1 detik
	
	for time := range channel {
		fmt.Println(time)
	}
}

