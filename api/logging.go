package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Log(h http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.New(os.Stdout, "[ROUTE] ", log.LstdFlags)
		start := time.Now()

		lrw := NewLoggingResponseWriter(w)
		h.ServeHTTP(lrw, r)

		logger.Printf(
			"%-5s%-40s\t%-30s\t%-10s\t%d\n",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
			lrw.statusCode,
		)
	})
}
