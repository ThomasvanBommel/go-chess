package main

import (
	"fmt"
	"errors"
)

type Board struct {
	pieces [8][8]Piece
	flipped bool
}

func (b *Board) init() {
	backrow := []string{ "Rook", "Knight", "Bishop", "Queen", "King", "Bishop", "Knight", "Rook" }

	for x, col := range b.pieces {
		for y, _ := range col {
			piece := ""

			// Backrow
			if y == 0 || y == 7 {
				piece = backrow[x]
			}

			// Pawns
			if y == 1 || y == 6 {
				piece = "Pawn"
			}

			if piece != "" {
				player := 1

				if y < 3 {
					player = 2
				}

				b.pieces[x][y] = pieceFactory(piece, player)
			}
		}
	}
}

func validatePosition(position [2]int) error {
	if position[0] < 0 || position[0] > 7 || position[1] < 0 || position[1] > 7 {
		return errors.New("Invalid position")
	}

	return nil
}

func (b Board) getPiece(position [2]int) (Piece, error) {
	err := validatePosition(position)
	if err != nil { return nil, err }

	return b.pieces[position[0]][position[1]], nil
}

func (b *Board) setPiece(position [2]int, piece Piece) error {
	err := validatePosition(position)
	if err != nil { return err }

	b.pieces[position[0]][position[1]] = piece

	return nil
}

func (b Board) print() {
	for i := 0; i < 3; i++ {
		if i == 1 {
			for y := 0; y < len(b.pieces); y++ {
				for x := 0; x < len(b.pieces[y]) + 2; x++ {
					if x == 0 || x == len(b.pieces[y]) + 1 {
						fmt.Print(8 - y)
					}else{
						if x - 1 == 0 {
							fmt.Print("|")
						}

						piece := b.pieces[x - 1][y]

						if piece != nil {
							fmt.Print(piece.getName(), piece.getPlayer(), "|")
						}else{
							fmt.Print("__|")
						}
					}
				}

				fmt.Println()
			}
		}else{
			fmt.Print(" ")

			for _, char := range []string{ "A", "B", "C", "D", "E", "F", "G", "H" } {
				fmt.Printf(" %v ", char)
			}

			fmt.Println();
		}
	}
}