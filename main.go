package main

import (
	"log"
	"net/http"
)

const ROOT = "/"

func home(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello from Snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Listening on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
