package main

import "fmt"

func main() {
	mahasiswa := Mahasiswa{Name: "rizki", Kelas: "01"}

	helloTanpaPointer(mahasiswa)
	fmt.Println(mahasiswa)

	helloPointer(&mahasiswa)
	fmt.Println(mahasiswa)
}

type Mahasiswa struct {
	Name, Kelas string
}

func helloTanpaPointer(mahaiswa Mahasiswa) {
	mahaiswa.Name = "mufrizal"
}

func helloPointer(mahaiswa *Mahasiswa) {
	mahaiswa.Name = "mufrizal"
}
