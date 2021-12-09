package main

type Queen struct {
	PieceBase
}

func (p Queen) validateMove(move Move, b *Board) (bool, error) {
	return false, nil
}