package config

import (
	"fmt"
	"log"
	"net/http"
)

// Hampir sama dengan cara ke 1, perbedaannya ada pada pengaturan variable dari setiap environtment

// deklarasi variable untuk setiap environment development dan production
const DEVELOPMENT_PORT = 3000
const DEVELOPMENT_APP_NAME = "Development"
const DEVELOPMENT_DEBUG_MODE = true
const PRODUCTION_PORT = 4000
const PRODUCTION_APP_NAME = "Production"
const PRODUCTION_DEBUG_MODE = false

type Environment struct {
	Port      int
	AppName   string
	DebugMode bool
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Berbeda dengan cara 1, cara 2 tidak menggunakan file json, jadi untuk konfigurasi environment nya berbeda
func LoadEnvironment(param string) (Environment, error) {
	env := Environment{ // memberikan nilai pada property2 pada struct Environment, variable ini berfungsi sebagi default jika environtment yg diinginkan tidak ada
		Port:      PRODUCTION_PORT,
		AppName:   PRODUCTION_APP_NAME,
		DebugMode: PRODUCTION_DEBUG_MODE,
	}

	switch param { // pengecekan environment
	case "PRODUCTION": // jika production maka akan tetap sesuai default
	case "DEVELOPMENT": // jika development maka nilai pada struct akan diubah sesuai ketentuan
		env.Port = DEVELOPMENT_PORT
		env.AppName = DEVELOPMENT_APP_NAME
		env.DebugMode = DEVELOPMENT_DEBUG_MODE
	default: // jika tidak ada akan diberikan output pemberitahuan dan environment akan di set default
		fmt.Println("Environment tidak ada")
	}

	return env, nil
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK} // fungsi yg merubah nilai dari struct loggingResponseWriter
}

func (conf Environment) WrapHandlerWithLogging() http.Handler { // method milik struct Environment yg berfungsi sebagai mereturn response dan status code

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		lrw := newLoggingResponseWriter(w) // memanggil fungsi newLoggingResponseWriter dengan memberi parameter w sebagai http.ResponseWriter

		statusCode := lrw.statusCode // deklrasi variable statusCode yg memiliki nilai dari hasil proses fungsi newLoggingResponseWriter

		text := req.URL.Path // variable yg membawa value Path Endpoint
		switch text {        // pengecekan endpoint apakah tersedia atau tidak
		case "/": // jika tersedia
			text = conf.AppName // maka akan memberikan nilai pada variable text sesuai yg sudah ditentukan
		case "/ping":
			text = "pong"
		case "/time":
			text = "waktu"
		default: // jika tidak tersedia, akan memberi return Not Found
			statusCode = http.StatusNotFound
			text = fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode))
		}

		fmt.Fprintf(w, text) // menampilkan response sesuai endpoint

		if conf.DebugMode { // jika debug mode adalah true
			fmt.Println(text) // maka akan menampilkan response
			// dan menampilkan method, path, status code, status text dari endpoint tersebut di terminal
			log.Printf("%s %s <-- %d %s", req.Method, req.URL.Path, statusCode, http.StatusText(statusCode))
		}
	})

}
