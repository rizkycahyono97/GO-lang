package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List =  list.New()

	// seperti struktur data queue
	data.PushBack("Rizky")
	data.PushBack("Cahyono")
	data.PushBack("Putra")

	var head = data.Front()	// cara memanggil yang paling depan
	fmt.Println(head.Value)	// rizky

	next := head.Next()
	fmt.Println(next.Value)	// cahyono

	// jika mengambil semua data menggunakan looping
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}


}