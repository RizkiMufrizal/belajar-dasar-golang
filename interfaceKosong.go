package main

import "fmt"

func main() {
	var data = Ups(1, "1")
	fmt.Println(data)
	fmt.Println(Ups(2, 2))
	fmt.Println(Ups(3, true))
}

func Ups(i int, sample interface{}) interface{} {
	fmt.Println(sample)
	if i == 1 {
		return i
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}
