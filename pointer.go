package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{City: "Jakarta", Country: "Indonesia", Province: "Jakarta"}
	address2 := &address1

	address1.City = "Aceh"

	fmt.Println(address1)
	fmt.Println(address2)

	address2.City = "Bali"

	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{City: "Malang", Country: "Indonesia", Province: "Jawa Timur"}

	fmt.Println(address1)
	fmt.Println(address2)

	//

	var address3 = new(Address)
	address3.City = "Jakarta"
	fmt.Println(address3)
}
