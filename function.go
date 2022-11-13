package main

import "fmt"

func main() {
	sayHello()
	sayHelloName("rizki", "mufrizal")
	fmt.Println(sayHelloReturn("mufrizal"))

	var names = []map[string]any{
		{"name": "rizki"},
		{"name": "mufrizal"},
	}

	fmt.Println(filterName(names, "rizki"))

	firstName, lastName, lengthName := sayHelloMultiName("rizki", "mufrizal")

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(lengthName)

	_, _, lengthOtherName := sayHelloMultiName("rizki", "mufrizal")
	fmt.Println(lengthOtherName)

	firstNameOther, _, _ := sayHelloNamedReturn()
	fmt.Println(firstNameOther)

	fmt.Println(sumAll(100, 10, 20, 30, 40, 50))
	fmt.Println(sumAll(100))

	numbers := []int{100, 200, 300}
	fmt.Println(sumAll(10, numbers...))

	goodBye := getGoodBye
	fmt.Println(goodBye("rizki"))
	fmt.Println(getGoodBye("mufrizal"))

	spamFIlter := spamFilter
	testSpamFilter := sayHelloFilter("spam", spamFIlter)
	fmt.Println(testSpamFilter)

	testNotSpam := sayHelloFilter("not spam", spamFIlter)
	fmt.Println(testNotSpam)

	fmt.Println(sayHelloFilterOther("not spam", spamFIlter))
}

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloName(fistName string, lastName string) {
	fmt.Println("hello", fistName, lastName)
}

// return data
func sayHelloReturn(name string) string {
	return "hello " + name
}

func filterName(names []map[string]any, name string) string {
	for _, n := range names {
		if n["name"] == name {
			return name
		}
	}
	return ""
}

// return multiple data
func sayHelloMultiName(firstName string, lastName string) (string, string, int) {
	return firstName, lastName, len(firstName + " " + lastName)
}

// return named return
func sayHelloNamedReturn() (firstName, lastName, middleName string) {
	firstName = "rizki"
	middleName = " "
	lastName = "mufrizal"
	return
}

// variadic function
func sumAll(otherNumber int, numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total + otherNumber
}

// function value
func getGoodBye(name string) string {
	return "good by " + name
}

// function as parameter
// function 1
func spamFilter(name string) string {
	if name == "spam" {
		return "this is spam"
	}
	return name
}

// function 2
func sayHelloFilter(name string, filterSpam func(string) string) string {
	return filterSpam(name)
}

type Filter func(string) string

func sayHelloFilterOther(name string, filter Filter) string {
	return filter(name)
}
