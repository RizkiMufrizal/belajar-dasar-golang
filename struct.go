package main

import "fmt"

func main() {
	var rizki Customer
	rizki.Name = "rizki"
	rizki.Address = "depok"
	rizki.Age = 20

	mufrizal := Customer{
		Name:    "mufrizal",
		Address: "jakarta",
		Age:     10,
	}

	riz := Customer{
		"riz",
		"bekasi",
		15,
	}

	fmt.Println(rizki)
	fmt.Println(rizki.Name)

	var data []Customer

	data = append(data, rizki)
	data = append(data, mufrizal)
	data = append(data, riz)

	fmt.Println(data)

	rizki.helloCustomer("joko")
}

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) helloCustomer(name string) {
	fmt.Println("hello", name, "My Name is", customer.Name)
}
