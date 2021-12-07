package main

import (
	"log"
	"net/http"
)

const ROOT = "/"
const NOT_ALLOWED_STATUS_CODE = 405

func home(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != ROOT {
		http.NotFound(writer, request)
		return
	}

	writer.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("showSnippet"))
}

func createSnippet(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		WriteNotAllowed(writer)
		return
	}

	writer.Write([]byte("createSnippet"))
}

func WriteNotAllowed(writer http.ResponseWriter) {
	writer.Header().Set("Allow", http.MethodPost)

	http.Error(writer, "Method Not Allowed", 405)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Listening on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
