package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != ROOT {
		http.NotFound(writer, request)
		return
	}

	filenames := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	template, error := template.ParseFiles(filenames...)
	if !checkErrorResponse(error, writer) {
		return
	}

	error = template.Execute(writer, nil)
	if !checkErrorResponse(error, writer) {
		return
	}
}

func showSnippet(writer http.ResponseWriter, request *http.Request) {
	idAsText := request.URL.Query().Get("id")
	id, err := strconv.Atoi(idAsText)

	if err != nil || id < 1 {
		writer.WriteHeader(404)
		fmt.Fprintf(writer, "Cannot handle ID '%s'", idAsText)
		return
	}

	fmt.Fprintf(writer, "Showing snippet with ID '%d'", id)
}

func createSnippet(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writeNotAllowed(writer)
		return
	}

	writer.Write([]byte("createSnippet"))
}

func writeNotAllowed(writer http.ResponseWriter) {
	writer.Header().Set("Allow", http.MethodPost)

	http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func checkErrorResponse(target error, writer http.ResponseWriter) bool {
	if target == nil {
		return true
	}
	log.Println(target.Error())
	http.Error(writer, "Internal server error", http.StatusInternalServerError)
	return false
}
