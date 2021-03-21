package data

type Square struct {
    Color   rune    `json:"color"`
    Piece   rune   `json:"piece"`
}

var Table [8][8]Square
var Moves []Move

func CreateTable() [8][8]Square {
    var finalSquare [8][8]Square
    for indx := 0; indx < 8; indx++ {
        for indy := 0; indy < 8; indy++ {
            if indy > 1 && indy < 6 {
                continue
            }
            if indy == 1 {
                finalSquare[indx][indy].Piece = 'P'
                finalSquare[indx][indy].Color = 'W'
                continue
            }
            if indy == 6 {
                finalSquare[indx][indy].Piece = 'P'
                finalSquare[indx][indy].Color = 'B'
                continue
            }
            if indy == 0 || indy == 7 {
                if indy == 0 {
                    finalSquare[indx][indy].Color = 'W'
                } else {
                    finalSquare[indx][indy].Color = 'B'
                }

                if indx == 0 || indx == 7 {
                    finalSquare[indx][indy].Piece = 'R'
                }
                if indx == 1 || indx == 6 {
                    finalSquare[indx][indy].Piece = 'N'
                }
                if indx == 2 || indx == 5 {
                    finalSquare[indx][indy].Piece = 'B'
                }
                if indx == 3 {
                    finalSquare[indx][indy].Piece = 'Q'
                }
                if indx == 4 {
                    finalSquare[indx][indy].Piece = 'K'
                }
            }
        }
    }
    return finalSquare
}