package web

import (
	"bytes"
	"encoding/json"
	"github.com/elvismdnin/matchManager/internal/data"
	"github.com/elvismdnin/matchManager/internal/web"
	"net/http"

	"github.com/gorilla/mux"
)

func CommandMove() web.Handler {
	return web.Handler {
		Route: func(r *mux.Route) {
			r.Path("/move").Methods("POST")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			var move data.Move
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(r.Body)
			body := buf.String()
			decodeErr := json.Unmarshal([]byte(body), &move)
			if decodeErr != nil {
				http.Error(w, decodeErr.Error(), http.StatusBadRequest)
				return
			}

			if !data.CheckValidMovement(move.OriX, move.OriY, move.DesX, move.DesY, data.Table) {
				http.Error(w, "Invalid movement", http.StatusBadRequest)
				return
			}

			data.Moves = append(data.Moves, move)
			data.Table = data.MovePiece(move, data.Table)

			encodeErr := json.NewEncoder(w).Encode(data.Table)
			if encodeErr != nil {
				http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
			}
		},
	}
}