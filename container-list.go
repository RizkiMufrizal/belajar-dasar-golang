package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("rizki")
	data.PushBack("mufrizal")
	data.PushBack("12345")
	data.PushFront("budi")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
