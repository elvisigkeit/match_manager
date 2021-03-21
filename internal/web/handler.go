package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Route func(r *mux.Route)
	Func  http.HandlerFunc
}

func (h Handler) AddRoute(r *mux.Router) {
	h.Route(r.NewRoute().HandlerFunc(h.Func))
}