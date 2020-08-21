package main

import (
    "github.com/elvismdnin/crius/web"
    "github.com/elvismdnin/crius/internal/data"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    data.AddInPlayers("1", "Archer", 2, 0)
    data.AddInPlayers("2", "Fighter", -2, 0)

    router := mux.NewRouter()
    router.HandleFunc("/players", web.GetPlayers).Methods("GET")
    router.HandleFunc("/position/{id}", web.GetPosition).Methods("GET")
	router.HandleFunc("/move/{id}", web.CommandMove).Methods("POST")
	router.HandleFunc("/skill/{id}", web.CommandSkill).Methods("POST")
	
    log.Fatal(http.ListenAndServe(":8000", router))
}