package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Hostname")
	number := flag.Int("number", 100, "Number")

	flag.Parse()

	fmt.Println("hostnane", *host)
	fmt.Println("Number", *number)
}
