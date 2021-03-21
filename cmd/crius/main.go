package main

import (
	"github.com/elvismdnin/crius/internal/data"
	"github.com/elvismdnin/crius/web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
    router := mux.NewRouter()
    data.Table = data.CreateTable()

    web.GetTable(data.Table).AddRoute(router)
    web.CommandMove().AddRoute(router)
	// // router.HandleFunc("/move/{id}", web.CommandMove).Methods("POST")
	// // router.HandleFunc("/skill/{id}", web.CommandSkill).Methods("POST")
	
    log.Fatal(http.ListenAndServe(":8000", router))
}