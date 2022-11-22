package main

import "fmt"

func main() {
	runApp(true)
	fmt.Println("test panic dan recover")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("aplikasi error")
	}
}

func endApp() {
	fmt.Println("aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("error dengan message", message)
	}
}
