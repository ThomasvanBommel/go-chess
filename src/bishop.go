package main

type Bishop struct {
	PieceBase
}

func (p Bishop) validateMove(move Move, b *Board) (bool, error) {
	return false, nil
}