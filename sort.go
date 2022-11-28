package main

import (
	"fmt"
	"sort"
)

func main() {
	users := []User{
		{Name: "rizki", Age: 30},
		{Name: "Mufrizal", Age: 15},
		{Name: "Budi", Age: 50},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}
