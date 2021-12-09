package main

var id = 0

func newPiece(name string, player int) Piece {
	id++

	switch name {
		case "Rook":
			return Rook{ PieceBase{ "R", player, id } }

		case "Knight":
			return Knight{ PieceBase{ "k", player, id } }

		case "Bishop":
			return Bishop{ PieceBase{ "B", player, id } }

		case "Queen":
			return Queen{ PieceBase{ "Q", player, id } }

		case "King":
			return King{ PieceBase{ "K", player, id } }

		default:
			return Pawn{ PieceBase{ "P", player, id }, false }
	}
}