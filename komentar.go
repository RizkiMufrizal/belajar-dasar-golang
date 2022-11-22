package main

import "fmt"

func main() {
	helloKalimat("rizki")

	kalimat := func(nama string) {
		fmt.Println("hallo", nama)
	}

	kalimat("rizki")
}

/*
* function ini untuk print parameter
 */
func helloKalimat(kalimat string) {
	fmt.Println("hello", kalimat)
}
