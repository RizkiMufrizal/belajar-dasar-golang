package test

import (
	"belajar-golang-restful-api/simple"
	"fmt"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	fmt.Println(err)
	fmt.Println(simpleService)
}
