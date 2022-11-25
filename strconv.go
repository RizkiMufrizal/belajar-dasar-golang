package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(boolean)

	iniString := strconv.FormatBool(false)
	fmt.Println(iniString)

	number, err := strconv.ParseInt("100000", 10, 32)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(number)

	value := strconv.FormatInt(number, 10)
	fmt.Println(value)
}
