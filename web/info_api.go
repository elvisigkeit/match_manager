package web

import (
    "github.com/elvismdnin/matchManager/internal/data"
    "github.com/elvismdnin/matchManager/internal/web"

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

//
//func GetPosition(players map[string]data.Player) web.Handler {
//    return web.Handler{
//        Route: func(r *mux.Route) {
//            r.Path("/position/{id}").Methods("GET")
//        },
//		Func: func(w http.ResponseWriter, r *http.Request) {
//            params := mux.Vars(r)
//            if player, ok := players[params["id"]]; ok {
//                pos := player.Pos
//                json.NewEncoder(w).Encode(pos)
//            } else {
//                log.Printf("ID doesn't correspond to a player\n")
//                http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
//            }
//        },
//    }
//}