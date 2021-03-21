package data

type Move struct {
	Color   rune    `json:"color"`
	OriX   int   `json:"originx"`
	OriY   int   `json:"originy"`
	DesX   int   `json:"destinx"`
	DesY   int   `json:"destiny"`
}

func CreateMove(color rune, orix int, oriy int, desx int, desy int) Move {
	var move Move

	move.Color = color
	move.OriX = orix
	move.DesX = desx
	move.OriY = oriy
	move.DesY = desy

	return move
}

func MovePiece(move Move, table [8][8]Square) [8][8]Square {
	table[move.DesX][move.DesY] = table[move.OriX][move.OriY]

	table[move.OriX][move.OriY].Color = 0
	table[move.OriX][move.OriY].Piece = 0

	return table
}