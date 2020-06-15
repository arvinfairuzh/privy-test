package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	key := []byte("1234privy5678")                          // key 13 byte
	plaintext := []byte("Selamat Datang")                   // yang akan diencrypt
	fmt.Printf("Text Sebelum di Encrypt : %s\n", plaintext) // menampilkan text yang belum diencrypt

	// saya membuat 2 cara untuk merubah key (32 byte dan 24 byte)
	nextKey := []byte("testarvinfairuzhuda") // merubah key menjadi 32 byte AES256
	str := append(key, nextKey...)           // merubah key menjadi 32 byte AES256
	// str := []byte(base32.StdEncoding.EncodeToString(key)) // merubah key menjadi 24 byte AES192

	ciphertext, err := encrypt(str, plaintext) // melakukan proses encrypt menggunakan fungsi encrypt
	if err != nil {                            // pengecekan error
		log.Fatal(err)
	}
	fmt.Printf("Hasil Encrypt : %0x\n", ciphertext) // menampilkan hasil encrypt

	result, err := decrypt(str, ciphertext) // melakukan proses decrypt menggunakan fungsi decrypt
	if err != nil {                         // pengecekan error
		log.Fatal(err)
	}
	fmt.Printf("Hasil Decrypt :  %s\n", result) // menampilkan hasil decrypt
}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) // membuat block cipher baru, menggunakan library cipher
	if err != nil {                  // jika err != nil
		return nil, err // akan menampilkan output error
	}
	ciphertext := make([]byte, aes.BlockSize+len(text)) // membuat array ciphertext
	iv := ciphertext[:aes.BlockSize]                    // membuat array iv yang memilik value sesuai dengan value ciphertext dimulai dari posisi awal value ciphertext sampai jumlah Block Size

	// pengecekan error
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// melakukan proses encrypt
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))

	// return hasil encrypt
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) // membuat block cipher baru, menggunakan library cipher
	if err != nil {                  // jika err != nil
		return nil, err // akan menampilkan output error
	}
	if len(text) < aes.BlockSize { // jika panjang variable text lebih kecil dari Block Size
		return nil, errors.New("ciphertext too short") // akan menampilkan output "ciphertext too short"
	}
	iv := text[:aes.BlockSize]  // membuat array iv yang memilik value sesuai dengan value text dimulai dari posisi awal value text sampai jumlah Block Size
	text = text[aes.BlockSize:] // memberikan value pada array text yang memiliki value sesuai dengan value text dimulai dari posisi yang ditentukan oleh jumlah Block Size sampai akhir value text

	// melakukan proses decrypt
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)

	// pengecekan error
	if err != nil {
		return nil, err
	}

	return text, nil
}
