package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arvinfairuzh/privy-test/soal-5/cara-1/config"
)

func main() {
	env, err := config.LoadEnvironment(config.ENVIRONMENT_DEVELOPMENT) // variable yg menggunakan fungsi pada package config
	if err != nil {                                                    // pengecekan error pada saat menggunakan fungsi LoadEnvironment
		log.Fatal(err)
	}

	rootHandler := env.WrapHandlerWithLogging() // memanggil method WrapHandlerWithLogging
	http.Handle("/", rootHandler)               // set route sesuai yg ditentukan
	http.Handle("/ping", rootHandler)
	http.Handle("/time", rootHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", env.Port), nil)) // menentukan port sesuai environment yg dipilih
}
