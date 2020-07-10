package api

import (
	"net/http"
	"github.com/gorilla/mux"
)

type api struct {
	router *mux.Router
}

func CreateAPI() *api {
	a := &api{
		router: mux.NewRouter().StrictSlash(true),
	}
	a.routes()
	return a
}

// Make api an http.Handler to make api and http.Handler usable interchangably
func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
