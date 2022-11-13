package main

import "fmt"

func main() {
	type NoKTP string

	var noKtp NoKTP = "1234567890"
	fmt.Println(noKtp)
}