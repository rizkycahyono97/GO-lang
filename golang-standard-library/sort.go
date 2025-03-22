package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}

type UserSlice []User 

// membuat kontrak untuk sort
func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Rizky", 30},
		{"Cahyono", 25},
		{"Putra", 20},
		{"Agus", 50},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}