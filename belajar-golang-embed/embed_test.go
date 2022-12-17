package belajargolangembed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestFileToString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}
