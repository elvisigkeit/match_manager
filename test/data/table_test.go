package data

import (
	"github.com/elvismdnin/crius/internal/data"
	
	"fmt"
	"testing"
)

func TestCreateTable(t *testing.T) {
	table := data.CreateTable()
	var strTable string
	for _, column := range table {
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
					"WR WP             BP BR "{
		t.Error("The table being created is wrong")
	}
}