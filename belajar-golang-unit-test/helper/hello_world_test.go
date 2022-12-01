package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")
	m.Run()
	fmt.Println("Setelah Unit Test")
}

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

func TestSKip(t *testing.T) {
	t.Log("Server", runtime.GOOS)
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on linux")
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Hello", func(t *testing.T) {
		result := HelloWorld("Rizki")
		assert.Equal(t, "Hello Rizki", result, "Result Must be 'Hello Rizki'")
	})
	t.Run("Hi", func(t *testing.T) {
		result := HiWorld("Rizki")
		require.Equal(t, "Hi Rizki", result, "Result Must be 'Hello Rizki'")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name      string
		request   string
		exptected string
	}{
		{
			name:      "HelloWorld(rizki)",
			request:   "Rizki",
			exptected: "Hello Rizki",
		},
		{
			name:      "HelloWorld(mufrizal)",
			request:   "Mufrizal",
			exptected: "Hello Mufrizal",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.exptected, result)
		})
	}
}
