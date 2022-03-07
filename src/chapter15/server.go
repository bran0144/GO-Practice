package chapter15

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "hello")
}
func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "bonjour")
}
func spanishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "hola")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/bonjour", frenchHandler)
	http.HandleFunc("/hola", spanishHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
