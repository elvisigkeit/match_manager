package main

import (
	"github.com/elvismdnin/match_manager/internal/data"
	"github.com/elvismdnin/match_manager/web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	port := "8000"
    router := mux.NewRouter()
    data.Table = data.CreateTable()

    web.GetTable(data.Table).AddRoute(router)
    web.CommandMove().AddRoute(router)

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server opened at " + port)
	log.Fatal(srv.ListenAndServe())
}