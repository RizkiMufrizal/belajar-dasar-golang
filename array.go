package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "rizki"
	names[1] = "mufrizal"
	names[2] = "rizki mufrizal"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(len(names))

	var values = [3]int{
		80, 90, 100,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}
