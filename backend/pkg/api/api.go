package api

import (
	"github.com/gorilla/mux"
	"net/http"

	"backend/pkg/middlewares"
)

type api struct {
	router *mux.Router
}

func CreateAPI() *api {
	a := &api{
		router: mux.NewRouter().StrictSlash(true),
	}
	a.router.Use(middlewares.Logger)
	a.routes()
	return a
}

// Make api an http.Handler to make api and http.Handler usable interchangably
func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
