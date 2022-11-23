package main

import "fmt"

func main() {
	rizki := Man{"Rizki"}
	rizki.Married()

	fmt.Println(rizki)

	rizki.MarriedPointer()
	fmt.Println(rizki)
}

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name
}
