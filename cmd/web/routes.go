package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (target *Application) routes(applicationConfig ApplicationConfig) http.Handler {
	standardMiddleware := alice.New(target.recoverPanic, target.logRequest, target.secureHeader)

	mux := http.NewServeMux()
	mux.HandleFunc("/", target.home)
	mux.HandleFunc("/snippet", target.showSnippet)
	mux.HandleFunc("/snippet/create", target.createSnippet)

	fileServer := http.FileServer(http.Dir(applicationConfig.StaticContentDirectory))
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
