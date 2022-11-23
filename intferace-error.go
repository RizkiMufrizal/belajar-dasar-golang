package main

import (
	"errors"
	"fmt"
)

func main() {
	var pembagian, err = pembagian(10, 5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(pembagian)
	}
}

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Error Pembagian")
	}
	return nilai / pembagi, nil
}
