package crius

import (
    "github.com/elvismdnin/crius/web"
    "github.com/elvismdnin/crius/internal/data"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    data.AddInPlayers("1", data.Player{Char: "Archer", Pos: data.Position{X: 2, Y: 0}})
    data.AddInPlayers("2", data.Player{Char: "Fighter", Pos: data.Position{X: -2, Y: 0}})

    fmt.Println(data.RetrievePlayers)

    router := mux.NewRouter()
    router.HandleFunc("/players", web.GetPlayers).Methods("GET")
    router.HandleFunc("/position/{id}", web.GetPosition).Methods("GET")
	router.HandleFunc("/move/{id}", web.CommandMove).Methods("POST")
	router.HandleFunc("/skill/{id}", web.CommandSkill).Methods("POST")
	
    log.Fatal(http.ListenAndServe(":8000", router))
}