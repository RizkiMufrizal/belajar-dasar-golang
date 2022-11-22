package main

import "fmt"

func main() {
	helloWorld(0)
}

func helloWorld(value int) {
	defer logging()
	fmt.Println("run hello world")
	result := 10 / value
	fmt.Println(result)
}

func logging() {
	fmt.Println("selesai memanggil function logging")
}
