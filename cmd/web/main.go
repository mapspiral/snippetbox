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

	app := &Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	server := &http.Server{
		Addr:     applicationConfig.Address,
		ErrorLog: errorLog,
		Handler:  app.routes(*applicationConfig),
	}

	infoLog.Printf("Starting server on %v", applicationConfig.Address)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
