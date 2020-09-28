package web

import (
    "github.com/elvismdnin/crius/internal/data"

    "encoding/json"
    "net/http"
    "log"

    "github.com/gorilla/mux"
)

func GetPlayers(players map[string]data.Player) data.Handler {
    return data.Handler{
        Route: func(r *mux.Route) {
            r.Path("/players").Methods("GET")
        },
		Func: func(w http.ResponseWriter, r *http.Request) {
            json.NewEncoder(w).Encode(players)
        },
    }
}

func GetPosition(players map[string]data.Player) data.Handler {
    return data.Handler{
        Route: func(r *mux.Route) {
            r.Path("/position/{id}").Methods("GET")
        },
		Func: func(w http.ResponseWriter, r *http.Request) {
            params := mux.Vars(r)
            if player, ok := players[params["id"]]; ok {
                pos := player.Pos
                json.NewEncoder(w).Encode(pos)
            } else {
                log.Printf("ID doesn't correspond to a player\n")
                http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
            }
        },
    }
}