package data

import (
	"github.com/elvismdnin/crius/internal/data"

	"testing"
)

func TestMovement(t *testing.T) {

}

func TestMovementWithoutPiece(t *testing.T) {
	table := data.CreateTable()

	if data.CheckValidMovement(0, 2, 0, 0, table) {
		t.Error("The movement needs a piece on its origin")
	}
}

func TestMovementTowardsOwnPiece(t *testing.T) {
	table := data.CreateTable()

	if data.CheckValidMovement(2, 0, 3, 0, table) {
		t.Error("Can't movement a piece towards it's own piece")
	}

	if !data.CheckValidMovement(2, 0, 2, 7, table) {
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