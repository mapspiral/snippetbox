package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

const ROOT = "/"

func main() {
	applicationConfig := new(ApplicationConfig)
	flag.StringVar(&applicationConfig.Address, "address", ":4000", "HTTP network adress")
	flag.StringVar(&applicationConfig.StaticContentDirectory, "statics", "./ui/static", "Static content directory")

	flag.Parse()

	infoLog := log.New(os.Stdin, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir(applicationConfig.StaticContentDirectory))
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := &http.Server{
		Addr:     applicationConfig.Address,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %v", applicationConfig.Address)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
