package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/mapspiral/snippetbox/pkg/models/mysql"
)

const ROOT = "/"

func main() {
	applicationConfig := new(ApplicationConfig)
	flag.StringVar(&applicationConfig.ConnectionString, "connectionString", "web:web@/snippetbox?parseTime=true", "Connection string")
	flag.StringVar(&applicationConfig.Address, "address", ":4000", "HTTP network address")
	flag.StringVar(&applicationConfig.StaticContentDirectory, "statics", "./ui/static", "Static content directory")

	flag.Parse()

	infoLog := log.New(os.Stdin, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, errorInfo := setupDatabase(applicationConfig.ConnectionString)

	if errorInfo != nil {
		errorLog.Fatal(errorInfo.Error())
	}

	defer db.Close()

	app := &Application{
		ErrorLog:     errorLog,
		InfoLog:      infoLog,
		SnippetModel: &mysql.SnippetModel{DB: db}}

	server := &http.Server{
		Addr:     applicationConfig.Address,
		ErrorLog: errorLog,
		Handler:  app.routes(*applicationConfig),
	}

	infoLog.Printf("Starting server on %v", applicationConfig.Address)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
