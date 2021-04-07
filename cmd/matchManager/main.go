package main

import (
	"github.com/elvismdnin/match_manager/internal/data"
	"github.com/elvismdnin/match_manager/web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := "8000"
    router := mux.NewRouter()
    data.Table = data.CreateTable()

    web.GetTable(data.Table).AddRoute(router)
    web.CommandMove().AddRoute(router)

	log.Println("Server opened at " + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}