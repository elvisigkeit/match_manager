package web

import (
    "github.com/elvismdnin/match_manager/internal/data"
    "github.com/elvismdnin/match_manager/internal/web"

    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
)

func GetTable(table [8][8]data.Square) web.Handler {
    return web.Handler {
        Route: func(r *mux.Route) {
            r.Path("/table").Methods("GET")
        },
		Func: func(w http.ResponseWriter, r *http.Request) {
            _ = json.NewEncoder(w).Encode(table)
        },
    }
}