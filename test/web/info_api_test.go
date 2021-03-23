package web

import (
	"fmt"
	"github.com/elvismdnin/crius/internal/data"
    "github.com/elvismdnin/crius/web"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetTable(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

    table := data.CreateTable()

    web.GetTable(table).AddRoute(r)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/table", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}

	var returnedTable [8][8]data.Square
	_ = json.Unmarshal(w.Body.Bytes(), &returnedTable)

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

	if strTable != 	"WR WP             BP BR "+
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