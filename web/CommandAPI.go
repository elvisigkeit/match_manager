package controller

import "encoding/json"

func CommandMove(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var pos Position
    _ = json.NewDecoder(r.Body).Decode(&pos)

    player := players[params["id"]]
    player.Pos.X = pos.X
    player.Pos.Y = pos.Y
    players[params["id"]] = player
    
    json.NewEncoder(w).Encode(players)
}

func CommandSkill(w http.ResponseWriter, r *http.Request) {
    
}