package main

import (
    "github.com/elvismdnin/crius/web"
    "github.com/elvismdnin/crius/internal/data"

    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    var players = make(map[string]Player)
    data.AddInPlayers(players, "1", "Archer", 2, 0)
    data.AddInPlayers(players, "2", "Fighter", -2, 0)

    router := mux.NewRouter()
    GetPlayers(players).AddRoute(router)
    GetPosition(players).AddRoute(router)    
	// router.HandleFunc("/move/{id}", web.CommandMove).Methods("POST")
	// router.HandleFunc("/skill/{id}", web.CommandSkill).Methods("POST")
	
    log.Fatal(http.ListenAndServe(":8000", router))
}