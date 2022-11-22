package main

import "fmt"

func main() {
	counter := 0
	name := "rizki"

	fmt.Println(counter)

	increment := func() {
		fmt.Println("increment")
		counter++
		name := "mufrizal"
		fmt.Println(name)
	}

	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
