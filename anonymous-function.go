package main

import "fmt"

type Blacklist func(string) bool

func RegisterUser(name string, blackList Blacklist) {
	if blackList(name) {
		fmt.Println("black list")
	} else {
		fmt.Println("not black list")
	}
}

func main() {
	blacklist := func(name string) bool { return name == "user1" }

	RegisterUser("user1", blacklist)
	RegisterUser("user2", func(name string) bool {
		return name == "user2"
	})

	RegisterUser("rizki", blacklist)
}
