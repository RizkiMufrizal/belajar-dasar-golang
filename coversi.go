package main

import "fmt"

func main() {
	var nilai32 int32 = 127
	var nilai64 = int64(nilai32)
	var nilai8 = int8(nilai32)

	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "rizki"
	var e = name[0]
	var eString = string(e)

	fmt.Println(eString)
}
