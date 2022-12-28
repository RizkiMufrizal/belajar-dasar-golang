package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncodeJson(t *testing.T) {
	logJson(1)
	logJson(true)
	logJson("rizki")
	logJson([]string{"data 1", "data 2", "data 3"})
	logJson(map[string]interface{}{
		"message": "success",
		"success": true,
	})
	logJson([]map[string]interface{}{
		{"productId": "1"},
		{"productId": "2"},
	})
}

type Barang struct {
	ProductId   string
	ProductName string
	Price       int
	Type        []string
	Details     []Detail
}

type Detail struct {
	Manufacture string
	Kode        string
}

func TestStructJson(t *testing.T) {
	barang := Barang{
		ProductId:   "12345",
		ProductName: "Rinso",
		Price:       1000,
		Type:        []string{"123", "456", "789"},
		Details: []Detail{
			{Manufacture: "Samsung", Kode: "A01"},
			{Manufacture: "IPhone", Kode: "A02"},
		},
	}
	logJson(barang)
}

type Pegawai struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestStructTag(t *testing.T) {
	pegawai := Pegawai{
		Name:     "rizki",
		ImageUrl: "https://rizkimufrizal.github.io/",
	}
	logJson(pegawai)

	pegawai1 := &Pegawai{}
	pegawaiData := `{"name":"rizki","image_url":"https://rizkimufrizal.github.io/"}`
	logDecodeJson(pegawaiData, pegawai1)
	fmt.Println(pegawai1)
}

func TestEncodeStream(t *testing.T) {
	writer, _ := os.Create("sample_encode.json")
	encoder := json.NewEncoder(writer)

	pegawai := Pegawai{
		Name:     "rizki",
		ImageUrl: "https://rizkimufrizal.github.io/",
	}
	_ = encoder.Encode(pegawai)
}
