package main

type Rook struct {
	PieceBase
}

func (p Rook) validateMove(move Move, b *Board) (bool, error) {
	return false, nil
}