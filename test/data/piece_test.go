package data

import (
	"github.com/elvismdnin/crius/internal/data"

	"testing"
)

func TestKingMovement(t *testing.T) {
	table := data.CreateTable()

	table[4][1].Piece = 0
	table[4][1].Color = 0

	if !data.CheckValidMovement(4, 0, 4, 1, table) {
		t.Error("The king should be able to move forward")
	}
	if data.CheckValidMovement(4, 0, 4, 2, table) {
		t.Error("The king shouldn't be able to move forward twice")
	}
}

func TestPawnMovement(t *testing.T) {
	table := data.CreateTable()

	if !data.CheckValidMovement(0, 1, 0, 2, table) {
		t.Error("The pawn should be able to move forward")
	}

	if !data.CheckValidMovement(0, 1, 0, 3, table) {
		t.Error("The pawn should be able to move twice at beginning")
	}

	if data.CheckValidMovement(0, 1, 0, 4, table) {
		t.Error("The pawn shouldn't be able to move forward thrice")
	}

	table[0][2].Piece = 'P'
	table[0][2].Color= 'W'

	table[1][3].Piece = 'P'
	table[1][3].Color= 'B'

	if data.CheckValidMovement(0, 2, 2, 2, table) {
		t.Error("The pawn shouldn't be able move horizontally")
	}

	if data.CheckValidMovement(0, 2, 0, 4, table) {
		t.Error("The pawn shouldn't be able to move twice at any moment")
	}

	if !data.CheckValidMovement(0, 2, 1, 3, table) {
		t.Error("The white pawn should be able to eat horizontally")
	}

	if !data.CheckValidMovement(1, 3, 0, 2, table) {
		t.Error("The black pawn should be able to eat horizontally")
	}

	if data.CheckValidMovement(0, 1, 1, 2, table) {
		t.Error("The pawn shouldn't be able to move diagonally")
	}
}

func TestKnightMovement(t *testing.T) {
	table := data.CreateTable()

	if !data.CheckValidMovement(1, 0, 2, 2, table) {
		t.Error("The knight should be able to move to right")
	}

	if !data.CheckValidMovement(1, 0, 0, 2, table) {
		t.Error("The knight should be able to move to left")
	}

	if data.CheckValidMovement(1, 0, 1, 2, table) {
		t.Error("The knight shouldn't be able to move forward twice")
	}
}

func TestBishopMovement(t *testing.T) {
	table := data.CreateTable()

	table[1][1].Piece = 0
	table[1][1].Color = 0
	table[2][1].Piece = 0
	table[2][1].Color = 0
	table[3][1].Piece = 0
	table[3][1].Color = 0

	if !data.CheckValidMovement(2, 0, 4, 2, table) {
		t.Error("The bishop should be able to move to right")
	}

	if !data.CheckValidMovement(2, 0, 0, 2, table) {
		t.Error("The bishop should be able to move to left")
	}

	if data.CheckValidMovement(2, 0, 2, 2, table) {
		t.Error("The bishop shouldn't be able to move forward only")
	}
}

func TestRookMovement(t *testing.T) {
	table := data.CreateTable()

	table[0][1].Piece = 0
	table[0][1].Color = 0
	table[1][1].Piece = 0
	table[1][1].Color = 0

	if !data.CheckValidMovement(0, 0, 0, 4, table) {
		t.Error("The rook should be able to move")
	}

	if data.CheckValidMovement(0, 0, 2, 2, table) {
		t.Error("The rook shouldn't be able to move diagonally")
	}
}

func TestQueenMovement(t *testing.T) {
	table := data.CreateTable()

	table[2][1].Piece = 0
	table[2][1].Color = 0
	table[3][1].Piece = 0
	table[3][1].Color = 0
	table[4][1].Piece = 0
	table[4][1].Color = 0

	if !data.CheckValidMovement(3, 0, 3, 4, table) {
		t.Error("The queen should be able to move like a rook")
	}

	if !data.CheckValidMovement(3, 0, 5, 2, table) {
		t.Error("The queen should be able to move like a bishop")
	}

	if data.CheckValidMovement(3, 0, 4, 2, table) {
		t.Error("The queen shouldn't be able to move like knights")
	}
}

func TestQueenMovementPassThrough(t *testing.T) {
	table := data.CreateTable()

	if !data.InvalidPassThrough(3, 0, 3, 2, table) {
		t.Error("This pass through of the front pawn is invalid")
	}

	if !data.InvalidPassThrough(3, 0, 5, 2, table) {
		t.Error("This pass through of the right pawn is invalid")
	}

	table[3][1].Piece = 0
	table[4][1].Piece = 0

	if data.InvalidPassThrough(3, 0, 3, 2, table) {
		t.Error("Without the front pawn, the movement is valid")
	}

	if data.InvalidPassThrough(3, 0, 5, 2, table) {
		t.Error("Without the right pawn, the movement is valid")
	}
}

func TestPawnMovementPassThrough(t *testing.T) {
	table := data.CreateTable()

	if data.InvalidPassThrough(0, 1, 0, 3, table) {
		t.Error("Without any obstacle it is valid")
	}

	table[0][2].Piece = 'N'
	table[0][2].Color = 'B'

	if !data.InvalidPassThrough(0, 1, 0, 3, table) {
		t.Error("This pass through of the knight is invalid")
	}
}

func TestRookMovementPassThrough(t *testing.T) {
	table := data.CreateTable()

	if !data.InvalidPassThrough(0, 0, 0, 2, table) {
		t.Error("This pass through of the pawn is invalid")
	}

	table[0][1].Piece = 0

	if data.InvalidPassThrough(0, 0, 0, 2, table) {
		t.Error("Without the pawn, the movement is valid")
	}
}

func TestBishopMovementPassThrough(t *testing.T) {
	table := data.CreateTable()

	if !data.InvalidPassThrough(2, 0, 4, 2, table) {
		t.Error("This pass through of the right pawn is invalid")
	}

	if !data.InvalidPassThrough(2, 0, 0, 2, table) {
		t.Error("This pass through of the left pawn is invalid")
	}

	table[3][1].Piece = 0
	table[1][1].Piece = 0

	if data.InvalidPassThrough(2, 0, 4, 2, table) {
		t.Error("Without the right pawn, the movement is valid")
	}

	if data.InvalidPassThrough(2, 0, 0, 2, table) {
		t.Error("Without the left pawn, the movement is valid")
	}
}

func TestMovementWithoutPiece(t *testing.T) {
	table := data.CreateTable()

	if data.CheckValidMovement(0, 2, 0, 0, table) {
		t.Error("The movement needs a piece on its origin")
	}
}

func TestMovementTowardsOwnPiece(t *testing.T) {
	table := data.CreateTable()

	if data.CheckValidMovement(0, 0, 0, 1, table) {
		t.Error("Can't movement a piece towards it's own piece")
	}

	table[0][1].Piece = 0
	table[0][1].Color = 0

	if !data.CheckValidMovement(0, 0, 0, 6, table) {
		t.Error("Movement should be possible toward another player's piece")
	}
}

func TestOutOfRangeMovement(t *testing.T) {
	table := data.CreateTable()

	if data.CheckValidMovement(-1, 0, 0, 0, table) {
		t.Error("The origin x shouldn't accept negative values")
	}
	if data.CheckValidMovement(0, -1, 0, 0, table) {
		t.Error("The origin y shouldn't accept negative values")
	}
	if data.CheckValidMovement(0, 0, -1, 0, table) {
		t.Error("The destination x shouldn't accept negative values")
	}
	if data.CheckValidMovement(0, 0, 0, -1, table) {
		t.Error("The destination y shouldn't accept negative values")
	}

	if data.CheckValidMovement(8, 0, 0, 0, table) {
		t.Error("The origin x shouldn't accept out of table values")
	}
	if data.CheckValidMovement(0, 8, 0, 0, table) {
		t.Error("The origin y shouldn't accept out of table values")
	}
	if data.CheckValidMovement(0, 0, 8, 0, table) {
		t.Error("The destination x shouldn't accept out of table values")
	}
	if data.CheckValidMovement(0, 0, 0, 8, table) {
		t.Error("The destination y shouldn't accept out of table values")
	}
}