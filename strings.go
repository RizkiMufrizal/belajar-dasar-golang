package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "   rizki mufrizal   "

	fmt.Println(strings.Trim(name, " "))
	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Contains(name, "mufrizal"))
	fmt.Println(strings.Split(name, " "))
	fmt.Println(strings.ReplaceAll(name, " ", ""))
}
