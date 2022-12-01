package helper

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Rizki")
	assert.Equal(t, "Hello Rizki", result, "Result Must be 'Hello Rizki'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Rizki")
	require.Equal(t, "Hello Rizki", result, "Result Must be 'Hello Rizki'")
}
