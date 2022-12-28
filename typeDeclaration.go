package main

import "fmt"

func main() {
	type NoKTP string

	var noKtp NoKTP = "1234567890"
	fmt.Println(noKtp)

	name := "rizki"
	nama := &name

	fmt.Println(name)
	fmt.Println(nama)
}

func sample(name *string)  {

}