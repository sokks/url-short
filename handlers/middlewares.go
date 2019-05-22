package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// PanicMiddleware recovers a panic is had so
func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// RequestLoggingMiddleware writes each request into a slice
func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("[ACCESS] ", r)
		reqString, _ := httputil.DumpRequest(r, true)
		log.Println("[ACCESS] ", r, reqString)

		next.ServeHTTP(w, r)
	})
}
