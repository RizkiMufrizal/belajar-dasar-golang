package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "rizki"

	if name == "rizki" {
		fmt.Println("hello " + name)
	} else if name == "mufrizal" {
		fmt.Println("hy " + name)
	} else {
		fmt.Println("hello other")
	}

	if name == "rizki" && len(name) == 5 {
		fmt.Println("Panjang " + strconv.Itoa(len(name)))
	}
	if len(name) < 10 {
		fmt.Println("other length " + name)
	}

	if length := len(name); length == 5 {
		fmt.Println("panjang sama dengan 5")
	} else if length > 10 {
		fmt.Println("lebih besar dari 10")
	}

}
