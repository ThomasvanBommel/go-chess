package main

type Piece interface {
	getName() string
	getPlayer() int
}

type PieceBase struct {
	name string
	player int
}

func (p PieceBase) getName() string {
	return p.name
}

func (p PieceBase) getPlayer() int {
	return p.player
}