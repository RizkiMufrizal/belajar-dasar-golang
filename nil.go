package main

import "fmt"

func main() {
	data := NewMap("")
	fmt.Println(NewMap("rizki"))
	fmt.Println(NewMap(""))

	if data == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println("data tidak kosong")
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}
