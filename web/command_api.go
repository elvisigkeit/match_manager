package web

import (
	"bytes"
	"encoding/json"
	"github.com/elvismdnin/match_manager/internal/data"
	"github.com/elvismdnin/match_manager/internal/web"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

func CommandMove() web.Handler {
	return web.Handler {
		Route: func(r *mux.Route) {
			r.Path("/move").Methods("POST")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			log.Println("Moving piece")
			var move data.Move
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(r.Body)
			body := buf.String()
			decodeErr := json.Unmarshal([]byte(body), &move)
			if decodeErr != nil {
				log.Println("Invalid move json")
				http.Error(w, decodeErr.Error(), http.StatusBadRequest)
				return
			}

			if !data.CheckValidMovement(move.OriX, move.OriY, move.DesX, move.DesY, data.Table) {
				log.Println("Invalid movement")
				http.Error(w, "Invalid movement", http.StatusBadRequest)
				return
			}

			data.Moves = append(data.Moves, move)
			data.Table = data.MovePiece(move, data.Table)

			encodeErr := json.NewEncoder(w).Encode(data.Table)
			if encodeErr != nil {
				log.Println("Problem with enconding")
				http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
			}
			log.Println("Success on moving")
		},
	}
}