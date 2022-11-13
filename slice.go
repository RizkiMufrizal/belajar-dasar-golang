package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	fmt.Println(slice1)

	slice1[0] = "Ubah Mei"
	fmt.Println(months)

	var slice2 = months[11:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Rizki")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"

	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	var newSlice = make([]string, 3, 5)
	newSlice[0] = "rizki"
	newSlice[1] = "mufrizal"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	initArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(initArray)
	fmt.Println(iniSlice)

	iniSlice = append(iniSlice, 10)

	fmt.Println(iniSlice)
}
