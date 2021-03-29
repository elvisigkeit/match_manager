package main

import (
	"github.com/elvismdnin/matchManager/internal/data"
	"github.com/elvismdnin/matchManager/web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
    router := mux.NewRouter()
    data.Table = data.CreateTable()

    web.GetTable(data.Table).AddRoute(router)
    web.CommandMove().AddRoute(router)
	
    log.Fatal(http.ListenAndServe(":8000", router))
}