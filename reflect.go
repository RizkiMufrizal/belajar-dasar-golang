package main

import (
	"fmt"
	"reflect"
)

func main() {
	pegawai := Pegawai{
		Age: 100,
	}

	pegawaiType := reflect.TypeOf(pegawai)
	fmt.Println(pegawaiType.NumField())
	fmt.Println(pegawaiType.Field(0).Name)
	fmt.Println(pegawaiType.Field(0).Tag.Get("required"))
	fmt.Println(pegawaiType.Field(0).Tag.Get("max"))

	fmt.Println(isValid(pegawai))
	pegawai.Age = 0
	fmt.Println(isValid(pegawai))
}

type Pegawai struct {
	Age int `required:"true" max:"10"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != 0
		}
	}
	return true
}
