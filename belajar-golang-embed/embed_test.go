package belajargolangembed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
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

//go:embed picture1.jpg
var picture []byte

func TestFileToByteArray(t *testing.T) {
	err := ioutil.WriteFile("picture2.jpg", picture, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
