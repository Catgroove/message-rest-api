package middlewares_test

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/middlewares"
)


func TestLogger(t *testing.T) {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(middlewares.Logger)
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte("test"))
    }).Methods("GET")

	ts := httptest.NewServer(r)
	defer ts.Close()

	http.Get(ts.URL)
}
