package middlewares

import (
	"net/http"
	"log"
)

// More for demonstration than being useful, at the moment
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}