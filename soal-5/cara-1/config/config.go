package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	DEFAULT_ENVIRONMENT_FILE = "config/environment.json" // variable untuk menyimpan path dari file environment.json

	ENVIRONMENT_DEVELOPMENT = "DEVELOPMENT" // variable yg digunakan untuk masuk ke environment development
	ENVIRONMENT_PRODUCTION  = "PRODUCTION"  // variable yg digunakan untuk masuk ke environment production
)

type Config struct {
	ListEnv []Environment `json:"CONFIG"` // struct yg menyimpan value dari key CONFIG pada file json, isimpan dalam bentuk slice yg memiliki nilai dari struct Environment
}

type Environment struct { // struct yg menyimpan value sesuai tag, tag tersebut digunakan untuk mapping informasi json ke property yg bersangkutan
	Environment string `json:"ENVIRONMENT"` // property Environment akan menampung data json property ENVIRONMENT
	Port        int    `json:"PORT"`
	AppName     string `json:"APP_NAME"`
	DebugMode   bool   `json:"DEBUG_MODE"`
}

type loggingResponseWriter struct { // struct yg menyimpan value response dan status code
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK} // fungsi yg merubah nilai dari struct loggingResponseWriter
}

func LoadEnvironment(env string) (Environment, error) { // funsi untuk koneksi antara struct dengan file json
	data := Config{}                                       // deklarasi variable yg memiliki nilai dari struct Config
	file, err := ioutil.ReadFile(DEFAULT_ENVIRONMENT_FILE) // variable yg membawa nilai dari file json
	if err != nil {                                        // pengecekan error pada saat mengambil data pada file json
		return Environment{}, err
	}

	if err := json.Unmarshal([]byte(file), &data); err != nil { // mengisi nilai property2 dari struct sesuai dengan data json property pada file json
		return Environment{}, err // dan juga dilakukan pengecekan error saat penggabungan, struct environment akan diberi nilai kosong
	}

	for _, v := range data.ListEnv { // pengecekan apakah environment yg diinginkan tersedia atau tidak pada list environment
		if v.Environment == env { // jika tersedia
			return v, nil // maka fungsi akan langsung mereturn nilai pada variable yg ada pada environtment yg diingkan ke dalam struct Environment
		}
	}

	fmt.Printf("environment '%s' tidak terdaftar, mengambil default environment...", env) // jika tidak tersedia

	return data.ListEnv[0], nil // maka akan mereturn data pada default environment (default di set sebagai Development)
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
