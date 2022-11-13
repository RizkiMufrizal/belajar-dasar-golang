package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "rizki",
		"address": "aceh",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	fmt.Println(len(person))

	delete(person, "title")

	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "belajar golang"
	book["name"] = "rizki"
	book["ups"] = "salah"
	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)

	//slice with map
	var mapSlice = []map[string]string{
		{"title": "belajar golang yuks"},
		{"name": "mufrizal"},
	}

	fmt.Println(mapSlice)
}
