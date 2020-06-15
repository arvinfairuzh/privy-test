package main

import (
	"fmt"
	"strings"
)

func main() {
	pertama := "19374048" // Deklarasi variable string pertama
	kedua := "aiueobcd"   // Deklarasi variable string kedua

	// Soal Nomer 1

	// Cara mengambil posisi angka ke 7 dari string pertama
	fmt.Println(strings.Index(pertama, "7")) // versi library golang

	// Versi Manual
	hasilPertama := 0           // Deklarasi variable hasilPertama, agar bisa diakses diluar for
	for i, v := range pertama { // looping variable pertama sesuai total karakter, i untuk menunjukan index array, v untuk menunjukan value
		if string(v) == "7" { // jika value sama dengan 7
			hasilPertama = i // maka akan menyimpan lokasi index dari value 7
		}
	}
	// Hasil Pertama adalah 3

	// Cara mengambil karakter mulai dari posisi ke 3 sampai akhir

	// Cara ke 1
	fmt.Println(string(kedua[hasilPertama:])) // menampilkan output dimulai dari posisi sesuai value hasilPertama, sampai dengan akhir posisi (:) pada variable kedua

	// Cara ke 2
	for i := hasilPertama; i < len(kedua); i++ { // looping variable kedua sesuai total karakter, dimulai dari index ke 3 sampai total karakter variable kedua
		fmt.Print(string(kedua[i])) // output looping variable kedua
	}

	// Hasilnya eobcd
}
