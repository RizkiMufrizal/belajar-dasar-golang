package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("argunment", args)

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(hostname)

	golangHome := os.Getenv("GO_HOME")
	fmt.Println("GO_HOME", golangHome)
}
