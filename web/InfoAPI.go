package web

import (
    "github.com/elvismdnin/crius/internal/data"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

func GetPlayers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(data.RetrievePlayers())
}

func GetPosition(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    players := data.RetrievePlayers()
    pos := players[params["id"]].Pos
    json.NewEncoder(w).Encode(pos)
}