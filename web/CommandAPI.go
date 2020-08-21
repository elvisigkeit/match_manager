package web

import (
    "github.com/elvismdnin/crius/internal/data"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

func CommandMove(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var pos data.Position
    _ = json.NewDecoder(r.Body).Decode(&pos)
    data.ChangePlayerPosition(params["id"], pos.X, pos.Y)
    players := data.RetrievePlayers()
    json.NewEncoder(w).Encode(players)
}

func CommandSkill(w http.ResponseWriter, r *http.Request) {
    
}