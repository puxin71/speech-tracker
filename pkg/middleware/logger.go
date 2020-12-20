package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := StatusWriter{ResponseWriter: w}

		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(&sw, r)

		log.Println(
			"Host:", r.Host,
			"Method:", r.Method,
			"RequestURI:", r.RemoteAddr,
			"Status:", http.StatusText(sw.statusCode),
			"Duration:", time.Since(start),
		)
	})
}

type StatusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *StatusWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}
