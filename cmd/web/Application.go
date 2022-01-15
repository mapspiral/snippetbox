package main

import (
	"log"

	"github.com/mapspiral/snippetbox/pkg/models/mysql"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger

	SnippetModel *mysql.SnippetModel
}
