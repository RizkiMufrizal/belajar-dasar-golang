package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func logDecodeJson(data string, parameter interface{}) {
	jsonByte := []byte(data)
	json.Unmarshal(jsonByte, parameter)
}

func TestDecodeJson(t *testing.T) {
	dataJson := `{"ProductId":"12345","ProductName":"Rinso","Price":1000}`
	product := &Barang{}
	logDecodeJson(dataJson, product)
	fmt.Println(product)
}

func TestDecodeToMap(t *testing.T) {
	data := `{"name":"rizki","image_url":"https://rizkimufrizal.github.io/"}`
	var penampung map[string]interface{}

	json.Unmarshal([]byte(data), &penampung)
	fmt.Println(penampung["name"])
	fmt.Println(penampung["image_url"])
}

func TestDecodeStream(t *testing.T) {
	reader, _ := os.Open("sample.json")
	decoder := json.NewDecoder(reader)

	pegawai := &Pegawai{}
	_ = decoder.Decode(pegawai)

	fmt.Println(pegawai)
}
