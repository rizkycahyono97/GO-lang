package golanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// =============== //
// channel = untuk menerima dan mengirim data Goroutines
// =============== //
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)	// hanya bisa mengirimkan 1 tipe data

	// channel <- "Rizky"	                                                                       // mengirim data dari data ke channel
	// data := <- channel	// menerima data dari channel ke data
	// fmt.Println(<- channel)	// langsung mengirim ke parameter

	defer close(channel) 	// harus di close supaya tidak memory leak	

	// test goroutine anonymous func 
	go func ()  {
		time.Sleep(2 * time.Second)
		channel <- "Rizky Cahyono Putra"	// mengirim data dari channel (wajib)
		fmt.Println("Selesai mengirim Data ke Channel")
	}()

	data := <- channel	// menerima data channer (wajib)
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}


// =========== //
// channel sebagai parameter
// =========== //
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Rizky Cahyono Putra"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}


// ============ //
// channel only in and out
// ============ //

// hanya mengirim channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Rizky Cahyono Putra"
}

// hanya menerima channel
func OnlyOut(channel <-chan string) {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}


// ============== //
// buffered channel
// ============== //
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)	// 3 -> buffered
	defer close(channel)

	go func() {
		channel <- "Rizky"
		channel <- "Cahyono"
		channel <- "Putra"
	}()

	go func () {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}


// ============== //
// range channel
// ============== //
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	// memasukan data ke channel
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}

		close(channel) // wajib	
	}()

	// menerima data dari loop channel diatas
	for data := range channel {
		fmt.Println("Menerima Data ", data)
	}

	fmt.Println("Selesai")
}


// ============== //
// select multiple channel
// ============== //
func TestSelectMultipleChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

// ============== //
// default select
// ============== //
func TestDefaultSelectMultipleChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}