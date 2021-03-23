package web

import (
	"encoding/json"
	"fmt"
	"github.com/elvismdnin/crius/internal/data"
	"github.com/elvismdnin/crius/web"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCommandMove(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	web.CommandMove().AddRoute(r)

	data.Table = data.CreateTable()
	data.Table[0][1].Color = 0
	data.Table[0][1].Piece = 0

	move := data.CreateMove('W', 0, 0, 0, 4)
	moveJson, _ := json.Marshal(move)
	moveBody := strings.NewReader(string(moveJson))

	r.ServeHTTP(w, httptest.NewRequest("POST", "/move", moveBody))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}

	var returnedTable [8][8]data.Square
	_ = json.Unmarshal([]byte(w.Body.String()), &returnedTable)

	var strTable string
	for _, column := range returnedTable {
		for _, val := range column {
			if val.Piece != 0 {
				square := fmt.Sprintf("%c%c ", val.Color,val.Piece)
				strTable += square
			} else {
				strTable += "   "
			}
		}
	}

	if strTable != 	"            WR    BP BR "+
		"WN WP             BP BN "+
		"WB WP             BP BB "+
		"WQ WP             BP BQ "+
		"WK WP             BP BK "+
		"WB WP             BP BB "+
		"WN WP             BP BN "+
		"WR WP             BP BR " {
		t.Error("The table being created is wrong")
	}
}