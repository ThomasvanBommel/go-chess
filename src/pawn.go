package main

import (
	"errors"
)

type Pawn struct {
	PieceBase
	hasMoved bool
}

func (p Pawn) validateMove(move Move, b *Board) (bool, error) {
	if move.distance[0] == 0 {
		// Didn't move anywhere ?!
		if move.distance[1] == 0 {
			return false, errors.New("You didn't move!")
		}

		// First move (2 cells forward)
		if !p.hasMoved && move.distance[1] == 2 {
			return true, nil
		}

		// Forward / Backward 1 cell
		if move.distance[1] == 1 {
			return true, nil
		}
	}

	// Attack (diagonal)
	if b.pieces[move.to[0]][move.to[1]] != nil && move.distance[0] == 1 && move.distance[1] == 1 {
		return true, nil
	}

	return false, errors.New("Did not pass validation.")
}