package main

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