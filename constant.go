package main

import "fmt"

func main() {
	const firstName = "rizki"
	const lastName = "mufrizal"
	const age = 30

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	const (
		otherName = "rizki mufrizal"
		otherAge  = 30
	)

	fmt.Println(otherName)
	fmt.Print(otherAge)
}
