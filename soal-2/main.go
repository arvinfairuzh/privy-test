package main

import (
	"fmt"
)

func cek(input string) (string, []string) { // Deklarasi Fungsi Cek dengan parameter input untuk melakukan pengecekan inputan

	vAwal := ""             // deklarasi variable vAwal sebagai string
	outputArr := []string{} // deklarasi array outputArr dengan value string

	for _, v := range input { // melakukan looping sesuai jumlah karakter variable input
		strOutput := "true" // deklarasi variable strOutput
		switch string(v) {  // jika string v adalah
		case "[", "{": // "[" atau "{"
			vAwal = string(v) // maka vAwal memiliki value string v
		case "]": // jika string v "]"
			if vAwal != "[" { // jika vAwal memiliki value bukan "["
				strOutput = "false" // maka strOutput memiliki value false
			}
			vAwal = "" // jika vAwal memiliki value "[", maka vAwal memilik value ""
		case "}": // jika string v "}"
			if vAwal != "{" { // jika vAwal memiliki value bukan "{"
				strOutput = "false" // maka strOutput memiliki value false
			}
			vAwal = "" // jika vAwal memiliki value "{", maka vAwal memilik value ""
		default: // jika string v bukan "[" , "{" , "]", atau "}"
			strOutput = "nil" //maka strOutput memiliki value nil
		}
		outputArr = append(outputArr, strOutput) // melakukan push ke array outputArr sesuai value dari strOutput
	}

	// jika vAwal ada isinya berarti ada inputan yg tidak memiliki pasangan, jadi outputnya langsung false
	if vAwal != "" { // jika vAwal tidak sama dengan "" atau ada isinya
		outputArr = append(outputArr, "false") // maka array outputArr memiliki value false
	}

	output := "true"              // deklarasikan variable output
	for _, v := range outputArr { // melakukan looping sesuai jumlah array outputArr
		switch v { // jika v
		case "nil": // adalah "nil"
			return v, outputArr // maka akan langsung melakukan return sesuai value v (nil) dan outputArr untuk melihat proses pengecekan
		case "false": // adalah "false"
			output = v // maka variable output akan memiliki value sesuai value v (false)
		}
	}

	return output, outputArr // melakukan return output dan outputArr untuk melihat proses pengecekan
}

func main() {
	output, outputArr := cek("asdasd") // memberikan input pada fungsi cek
	fmt.Println(output, outputArr)     // menampilkan value dari variable output
}

// Contoh input dan output
// [][]eqweqe -> nil
// [][] -> true
// [ -> false
// [][ -> false
// []{} -> true
// []{ -> false
// asdasd -> nil
