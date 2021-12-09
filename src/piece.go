package main

import "errors"

type Piece interface {
	getName() string
	getPlayer() int
	getID() int
	isMe(Piece) bool
	validateMove(Move, *Board) (bool, error)
}

type PieceBase struct {
	name string
	player int
	id int
}

func (p PieceBase) getName() string {
	return p.name
}

func (p PieceBase) getPlayer() int {
	return p.player
}

func (p PieceBase) getID() int {
	return p.id
}

func (p PieceBase) isMe(p2 Piece) bool {
	if p.id == p2.getID() {
		return true
	}

	return false
}

func (p PieceBase) validateMove(move Move, b *Board) (bool, error) {
	if !p.isMe(b.pieces[move.from[0]][move.from[1]]) {
		return false, errors.New("Wrong piece.")
	}

	return true, nil
}