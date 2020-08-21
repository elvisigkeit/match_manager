package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Position struct {
    X  int `json:"x"`
    Y  int `json:"y"`
}

type Player struct {
    Char 	  string   `json:"char"`
    Pos       Position `json:"pos"`
}

var players = make(map[string]Player)

func GetPlayers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(players)
}

func GetPosition(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    pos := players[params["id"]].Pos
    json.NewEncoder(w).Encode(pos)
}

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

func main() {
    players["1"] = Player{Char: "Archer", Pos: Position{X: 2, Y: 0}}
    players["2"] = Player{Char: "Fighter", Pos: Position{X: -2, Y: 0}}

    router := mux.NewRouter()
    router.HandleFunc("/players", GetPlayers).Methods("GET")
    router.HandleFunc("/position/{id}", GetPosition).Methods("GET")
	router.HandleFunc("/move/{id}", CommandMove).Methods("POST")
	router.HandleFunc("/skill/{id}", CommandSkill).Methods("POST")
	
    log.Fatal(http.ListenAndServe(":8000", router))
}