package main

type King struct {
	PieceBase
}

func (p King) validateMove(move Move, b *Board) (bool, error) {
	return false, nil
}