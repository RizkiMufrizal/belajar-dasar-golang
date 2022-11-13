package main

import "fmt"

func main() {
	var name string

	name = "rizki"
	fmt.Println(name)

	name = "mufrizal"
	fmt.Println(name)

	var otherName = "other rizki"
	fmt.Println(otherName)

	otherName = "other mufrizal"
	fmt.Println(otherName)

	var age int8 = 30
	fmt.Println(age)

	nameOther := "ok rizki"
	fmt.Println(nameOther)

	nameOther = "ok mufrizal"
	fmt.Println(nameOther)

	var (
		firstName = "rizki"
		lastName  = "mufrizal"
		ageOther  = 30
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(ageOther)
}
