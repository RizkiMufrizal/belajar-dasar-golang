package helper

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Rizki")
	if result != "Hello Rizki" {
		t.Error("Result Must be 'Hello Rizki'")
	}
	t.Log("Test Hello")
}

func TestHiWorld(t *testing.T) {
	result := HiWorld("Rizki")
	if result != "Hi Rizki" {
		t.Fatal("Result Must be 'Hi Rizki'")
	}
	t.Log("Test Hi")
}
