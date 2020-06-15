package main

import "fmt"

func fungsiA(slice []string) []string { // deklarasi fungsi fungsiA dengan parameter variable slice sebagai type slice string dan return array string
	fungsiMap := make(map[string]struct{}) // deklarasi map dengan key berupa string dan value berupa struct
	for _, v := range slice {              // looping sesuai length variable slice
		fungsiMap[v] = struct{}{} // mengisi fungsiMap dengan key berupa isi dari v, dengan value struct kosong
	}

	fungsiSlice := make([]string, 0, len(fungsiMap)) // deklarasi variable fungsiSlice, berupa slice string, length 0, capacity sesuai length fungsiMap
	for v := range fungsiMap {                       // looping sesuai length variable fungsiMap
		fungsiSlice = append(fungsiSlice, v) // melakukan push pada slice fungsiSlice, dengan value v yang merupakan key dari fungsiMap
	}
	return fungsiSlice // melakukan return value fungsiSlice
}

func main() {
	input := []string{"apple", "grape", "banana", "melon", "tahu", "tempe"} // mendeklarasikan slice string
	output := fungsiA(input)                                                // menyimpan value dari proses fungsiA ke dalam variable output
	fmt.Println(input)                                                      // menampilkan value variable input
	fmt.Println(output)                                                     // menampilkan value variable output
}
