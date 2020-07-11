package middlewares

import (
	"log"
	"net/http"
)

// More for demonstration than being useful, at the moment
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next.ServeHTTP(w, r)
	})
}
