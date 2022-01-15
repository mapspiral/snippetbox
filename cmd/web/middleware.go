package main

import (
	"fmt"
	"net/http"
)

func (target *Application) secureHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("X-XSS-Protection", "1; mode=block")
			writer.Header().Set("X-Frame-Options", "deny")

			next.ServeHTTP(writer, request)
		})
}

func (target *Application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			target.InfoLog.Printf("%s - %s %s %s", request.RemoteAddr, request.Proto, request.Method, request.URL.RequestURI())

			next.ServeHTTP(writer, request)
		})
}

func (target *Application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				writer.Header().Set("Connection", "close")
				target.serverError(writer, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(writer, request)
	})
}
