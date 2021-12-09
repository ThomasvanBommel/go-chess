package main

import (
	"errors"
)

type Pawn struct {
	PieceBase
	hasMoved bool
}

func (p Pawn) validateMove(move Move, b *Board) (bool, error) {
	// Piece / Player validation
	fromPiece, err := b.getPiece(move.from)

	if err != nil {
		return false, err
	}

	if !p.isMe(fromPiece) {
		return false, errors.New("Wrong piece.")
	}

	if p.getPlayer() != fromPiece.getPlayer() {
		return false, errors.New("That's not your piece!")
	}

	// Attacking
	attacking := false
	toPiece, err := b.getPiece(move.to)

	if err != nil {
		return false, err
	}

	if toPiece != nil && toPiece.getPlayer() == p.getPlayer() {
		attacking = true
	}

	if attacking && move.distance[0] != 1 && move.distance[1] != 1 {
		return false, errors.New("Invalid attack")
	}

	// Moving / First move
	if move.distance[0] > 1 || (p.hasMoved && move.distance[1] > 1) {
		return false, errors.New("Invalid move")
	}

	return true, nil
}