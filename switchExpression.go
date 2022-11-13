package main

import "fmt"

func main() {
	name := "rizki"

	switch name {
	case "rizki":
		fmt.Println("hello rizki")
		fmt.Println(len("rizki"))
	case "mufrizal":
		fmt.Println("hello mufrizal")
	default:
		fmt.Println("default gaes")
	}

	switch age := 100; age > 50 {
	case true:
		fmt.Println("lebih besar dari 50")
	case false:
		fmt.Println("kurang dari 50")
	}

	length := len(name)
	switch {
	case length > 50:
		fmt.Println("lebih dari 50")
	case length < 50:
		fmt.Println("kurang dari 50")
	default:
		fmt.Println("kosong gaes")
	}
}
