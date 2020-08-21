package crius

import (
    "github.com/elvismdnin/crius/controller"
    "fmt"
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

func retrievePlayers() map[string]Player{
    return players
}

func addInPlayers(id string, player Player){
    players[id] = player
}

func main() {
    addInPlayers("1", Player{Char: "Archer", Pos: Position{X: 2, Y: 0}})
    addInPlayers("2", Player{Char: "Fighter", Pos: Position{X: -2, Y: 0}})

    fmt.Println(retrievePlayers)

    router := mux.NewRouter()
    router.HandleFunc("/players", GetPlayers).Methods("GET")
    router.HandleFunc("/position/{id}", GetPosition).Methods("GET")
	router.HandleFunc("/move/{id}", CommandMove).Methods("POST")
	router.HandleFunc("/skill/{id}", CommandSkill).Methods("POST")
	
    log.Fatal(http.ListenAndServe(":8000", router))
}