package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("perulangan i ke", i)
	}

	contohSlice := []map[string]string{
		{"name": "rizki"},
		{"name": "mufrizal"},
	}

	for i := 0; i < len(contohSlice); i++ {
		fmt.Println(contohSlice[i])
		fmt.Println(contohSlice[i]["name"])
	}

	for i, m := range contohSlice {
		fmt.Println("index", i, "=", m["name"])
	}

	for _, m := range contohSlice {
		fmt.Println(m["name"])
	}

	contohMap := map[string]string{
		"first": "rizki",
		"last":  "mufrizal",
	}

	for _, s2 := range contohMap {
		fmt.Println(s2)
	}
}
