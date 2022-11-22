package main

import "fmt"

func main() {
	var person Person
	person.Name = "rizki"

	SayHello(person)

	cat := Animal{
		Name: "Meong",
	}

	SayHello(cat)
}

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
