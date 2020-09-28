package web

import (
	"github.com/elvismdnin/crius/internal/data"
    "github.com/elvismdnin/crius/web"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetPlayers(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

    var players = make(map[string]data.Player)
    data.AddInPlayers(players, "1", "Archer", 2, 0)

    web.GetPlayers(players).AddRoute(r)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/players", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}

	var returned_players map[string]data.Player
	json.Unmarshal([]byte(w.Body.String()), &returned_players)

	player := returned_players["1"]

	if player.Pos.X != 2 || player.Pos.Y != 0 {
		t.Error("The position from player must be (2, 0)")
	}
	if player.Char != "Archer" {
		t.Error("The player character must be Archer")
	}
}