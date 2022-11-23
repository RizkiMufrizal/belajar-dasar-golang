package main

import "fmt"

func main() {
	defer errorTypeAssertions()
	randomAja := random()
	randomString := randomAja.(string)
	fmt.Println(randomString)

	switch randomAja.(type) {
	case string:
		fmt.Println("string", randomAja)
	case int:
		fmt.Println("int", randomAja)
	}

	randomInt := randomAja.(int)
	fmt.Println(randomInt)
}

func random() interface{} {
	return "OK"
}

func errorTypeAssertions() {
	fmt.Println("error yang gaes")
}
