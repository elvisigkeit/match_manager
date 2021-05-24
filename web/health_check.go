package web

import (
    "fmt"
    "github.com/elvismdnin/match_manager/internal/web"

    "github.com/gorilla/mux"
    "net/http"
)

func HealthCheck() web.Handler {
    return web.Handler {
        Route: func(r *mux.Route) {
            r.Path("/").Methods("GET")
        },
		Func: func(w http.ResponseWriter, r *http.Request) {
            _, _ = fmt.Fprint(w, "{success: true}")
        },
    }
}