package main

import (
	"fmt"
)

func introduction() {
	fmt.Println("Welcome to go-chess!")
	fmt.Println()
}

func requestMove() *Move {
	fmt.Print("Enter your move: ")

	var input string
	fmt.Scanln(&input)

	move, err := stringToMove(input)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Try again...")
		return requestMove()
	}

	return move
}

var firstPlayersTurn = true

func turn(b *Board) bool {
	player := 1

	if !firstPlayersTurn {
		player = 2
	}

	fmt.Println()
	fmt.Println("Player #", player, ", your turn!")

	move := requestMove()
	// piece := b.pieces[move.from[0]][move.from[1]]
	piece, _ := b.getPiece(move.from) // missing err

	for {
		valid, err := piece.validateMove(*move, b)

		if err != nil {
			fmt.Println(err)
		}

		if valid {
			break
		}

		move = requestMove()
	}
	
	// fmt.Println("Piece:", piece)

	// b.pieces[move.to[0]][move.to[1]] = piece
	b.setPiece(move.to, piece)
	b.setPiece(move.from, nil)

	firstPlayersTurn = !firstPlayersTurn

	return false
}

func main() {
	introduction()

	board := Board{}
	board.init()
	board.print()

	for {
		turn(&board)
		board.print()
	}
}