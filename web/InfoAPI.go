package controller

import "encoding/json"

func GetPlayers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(players)
}

func GetPosition(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    pos := players[params["id"]].Pos
    json.NewEncoder(w).Encode(pos)
}