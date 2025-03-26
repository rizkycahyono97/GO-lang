package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

// =================== //
// map
// =================== //

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()
	
	// menampilkan isi dari sync.map
	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})

}