package main

type Knight struct {
	PieceBase
}

func (p Knight) validateMove(move Move, b *Board) (bool, error) {
	return false, nil
}