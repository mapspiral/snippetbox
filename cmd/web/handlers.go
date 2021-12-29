package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
	"strconv"
)

func (target *Application) home(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != ROOT {
		target.notFound(writer)
		return
	}

	filenames := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	template, errorInfo := template.ParseFiles(filenames...)
	if errorInfo != nil {
		target.serverError(writer, errorInfo)
	}

	errorInfo = template.Execute(writer, nil)
	if errorInfo != nil {
		target.serverError(writer, errorInfo)
	}
}

func (target *Application) showSnippet(writer http.ResponseWriter, request *http.Request) {
	idAsText := request.URL.Query().Get("id")
	id, errorInfo := strconv.Atoi(idAsText)

	if errorInfo != nil || id < 1 {
		writer.WriteHeader(404)
		fmt.Fprintf(writer, "Cannot handle ID '%s'", idAsText)
		return
	}

	fmt.Fprintf(writer, "Showing snippet with ID '%d'", id)
}

func (target *Application) createSnippet(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		target.writeNotAllowed(writer)
		return
	}

	writer.Write([]byte("createSnippet"))
}

func (target *Application) writeNotAllowed(writer http.ResponseWriter) {
	writer.Header().Set("Allow", http.MethodPost)

	target.clientError(writer, http.StatusMethodNotAllowed)
}

func (target *Application) serverError(writer http.ResponseWriter, errorInfo error) {
	trace := fmt.Sprintf("%s\n%s", errorInfo.Error(), debug.Stack())
	target.ErrorLog.Output(2, trace)
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (target *Application) clientError(writer http.ResponseWriter, statusCode int) {
	http.Error(writer, http.StatusText(statusCode), statusCode)
}

func (target *Application) notFound(writer http.ResponseWriter) {
	target.clientError(writer, http.StatusNotFound)
}
